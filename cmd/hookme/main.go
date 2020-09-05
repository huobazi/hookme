package main

import (
	"fmt"
	"github.com/huobazi/hookme/internal/config"
	"github.com/huobazi/hookme/internal/hooker"
	"github.com/huobazi/hookme/pkg/routes"
	"github.com/huobazi/hookme/pkg/voiceover"
	"net/http"
)

var (
	conf = config.Config
)

// Build version in compile-time
// see also https://github.com/ahmetb/govvv
var (
	GitCommit  string
	GitBranch  string
	GitState   string
	GitSummary string
	BuildDate  string
	Version    string
)

func hello(w http.ResponseWriter, _ *http.Request) {
	_, _ = fmt.Fprint(w, "Hookme server is starting ...\n")
}
func helloByName(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "Hello %s \n", routes.GetParam(r, 0))
}

func addRoute(hooker hooker.Hooker) {
	routes.AddRoute(hooker.GetRequestPath(), hooker.GetHttpMethods(), hooker.Hook)
}

func main() {

	addr := fmt.Sprintf("%s:%d", conf.Server.Host, conf.Server.Port)

	fmt.Println("Hookme server is starting ...")
	fmt.Printf("* Version %s\t\n", Version)
	fmt.Printf("* Build date %s\t\n", BuildDate)
	fmt.Printf("* Git branch %s\t\n", GitBranch)
	fmt.Printf("* Git summary %s\t\n", GitSummary)
	fmt.Printf("* Git commit %s\t\n", GitCommit)
	fmt.Printf("* Git state %s \t\n", GitState)
	fmt.Println("* Listening on tcp://", addr)
	fmt.Println("Use Ctrl-C to stop")

	router := routes.
		AddRoute("/hello", routes.MethodCollection{routes.GET}, hello).
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
		}else if cfg.Type == "docker_hub_hooker" {
			h := hooker.NewDockerHubHooker(hooker.BaseHooker{
				Name:        cfg.Name,
				RequestPath: cfg.RequestPath,
				WorkDir:     cfg.WorkDir,
				Command:     cfg.Command,
			})
			addRoute(h)
		}
	}

	err := http.ListenAndServe(addr, router)
	if err != nil {
		voiceover.Say("Hookme server running failed:", err)
	}
}