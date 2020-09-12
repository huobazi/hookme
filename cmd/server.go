package cmd

import (
	"context"
	"fmt"
	"github.com/huobazi/hookme/internal/config"
	"github.com/huobazi/hookme/internal/constants"
	"github.com/huobazi/hookme/internal/hooker"
	"github.com/huobazi/hookme/pkg/routes"
	"github.com/huobazi/hookme/pkg/voiceover"
	"github.com/sony/sonyflake"
	"github.com/spf13/cobra"
	"net/http"
	"os"
	"os/signal"
	"sync/atomic"
	"time"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Run hookme server",
	Long:  `This subcommand run hookme server`,
	Run: func(cmd *cobra.Command, args []string) {
		config.InitConfig(cfgFile)
		startServer()
	},
}

func init() {
	serverCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/hookme.yaml)")
	rootCmd.AddCommand(serverCmd)
}

var (
	conf    			= &config.Config
	snowflakeStartTime 	= time.Date(2020, 9, 1, 0, 0, 0, 0, time.UTC)
	healthy int32
)

type key int

const (
	requestIDKey key = 0
)

func hello(w http.ResponseWriter, _ *http.Request) {
	_, _ = fmt.Fprint(w, "Hookme server is starting ...\n")
}
func helloByName(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "Hello %s \n", routes.GetParam(r, 0))
}
func healthz(w http.ResponseWriter, _ *http.Request) {
	if atomic.LoadInt32(&healthy) == 1 {
		w.WriteHeader(http.StatusNoContent)
		return
	}
	w.WriteHeader(http.StatusServiceUnavailable)
}
func addRoute(hooker hooker.Hooker) {
	routes.AddRoute(hooker.GetRequestPath(), hooker.GetHttpMethods(), hooker.Hook)
}

func startServer() {
	addr := fmt.Sprintf("%s:%d", conf.Server.Host, conf.Server.Port)

	fmt.Println("Hookme server is starting ...")
	fmt.Printf("* Version %s\t\n", constants.Version)
	fmt.Printf("* Build date %s\t\n", constants.BuildDate)
	fmt.Printf("* Git branch %s\t\n", constants.GitBranch)
	fmt.Printf("* Git summary %s\t\n", constants.GitSummary)
	fmt.Printf("* Git commit %s\t\n", constants.GitCommit)
	fmt.Printf("* Git state %s \t\n", constants.GitState)
	fmt.Println("* Listening on ", addr)
	fmt.Println("Use Ctrl-C to stop")

	router := routes.
		AddRoute("/hello", routes.MethodCollection{routes.GET}, hello).
		AddRoute("/healthz", routes.MethodCollection{routes.GET}, healthz).
		AddRoute("/hello/([^/]+)", routes.MethodCollection{routes.GET}, helloByName)

	for _, cfg := range conf.Hooks {
		if cfg.Type == "aliyun_registry_hooker" {
			h := hooker.NewAliyunRegistryHooker(hooker.BaseHooker{
				Name:        cfg.Name,
				RequestPath: cfg.RequestPath,
				WorkDir:     cfg.WorkDir,
				Command:     cfg.Command,
			})
			addRoute(h)
		} else if cfg.Type == "docker_hub_hooker" {
			h := hooker.NewDockerHubHooker(hooker.BaseHooker{
				Name:        cfg.Name,
				RequestPath: cfg.RequestPath,
				WorkDir:     cfg.WorkDir,
				Command:     cfg.Command,
			})
			addRoute(h)
		}
	}

	sf := sonyflake.NewSonyflake(sonyflake.Settings{
		StartTime: snowflakeStartTime,
	})

	nextRequestID := func() string {
		if sf == nil{
			return fmt.Sprintf("%d", time.Now().UnixNano())
		}
		id, err := sf.NextID()
		if err != nil {
			return fmt.Sprintf("%d", time.Now().UnixNano())
		}
		return fmt.Sprintf("%d", id)
	}
	server := &http.Server{
		Addr:         addr,
		Handler:      tracing(nextRequestID)(logging()(router)),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	done := make(chan bool)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	go func() {
		<-quit
		voiceover.Say("Hookme server is shutting down...")
		atomic.StoreInt32(&healthy, 0)

		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		server.SetKeepAlivesEnabled(false)
		if err := server.Shutdown(ctx); err != nil {
			voiceover.Sayf("Hookme server could not gracefully shutdown: %v\n", err)
		}
		close(done)
	}()

	atomic.StoreInt32(&healthy, 1)

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		voiceover.Sayf("Hookme server could not listen on %s: %v\n", addr, err)
	}

	<-done
	voiceover.Say("Hookme server stopped")
}

func logging() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				requestID, ok := r.Context().Value(requestIDKey).(string)
				if !ok {
					requestID = "unknown"
				}
				voiceover.Say(requestID, r.Method, r.URL.Path, r.RemoteAddr, r.UserAgent())
			}()
			next.ServeHTTP(w, r)
		})
	}
}

func tracing(nextRequestID func() string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			requestID := r.Header.Get("X-Request-Id")
			if requestID == "" {
				requestID = nextRequestID()
			}
			ctx := context.WithValue(r.Context(), requestIDKey, requestID)
			w.Header().Set("X-Request-Id", requestID)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}