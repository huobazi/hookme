package main

import (
	"fmt"
	"github.com/huobazi/hookme/internal/config"
	"github.com/huobazi/hookme/internal/constants"
	"github.com/huobazi/hookme/internal/hooker"
	"github.com/huobazi/hookme/pkg/routes"
	"github.com/huobazi/hookme/pkg/voiceover"
	"net/http"
)

var (
	conf = config.Config
)

func hello(w http.ResponseWriter, _ *http.Request) {
	_, _ = fmt.Fprintf(w, "Hookme server %s is starting ...\n", constants.HookmeVersion)
}
func helloByName(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "Hookme server %s is starting ...\n", constants.HookmeVersion)
	_, _ = fmt.Fprintf(w, "Hello %s \n", routes.GetParam(r, 0))
}

func addRoute(hooker hooker.IHooker) {
	routes.Router.AddRoute(hooker.GetRequestPath(), hooker.GetHttpMethods(), hooker.Hook)
}

func main() {

	addr := fmt.Sprintf("%s:%d", conf.Server.Host, conf.Server.Port)

	fmt.Println("Hookme server is starting ...")
	fmt.Printf("* Version %s \n", constants.HookmeVersion)
	fmt.Println("* Listening on tcp://", addr)
	fmt.Println("Use Ctrl-C to stop")

	routes.Router.
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
		}
	}

	err := http.ListenAndServe(addr, routes.Router)
	if err != nil {
		voiceover.Say("Hookme server running failed:", err)
	}
}
