package hooker

import (
	"fmt"
	"github.com/huobazi/hookme/pkg/routes"
	"github.com/huobazi/hookme/pkg/voiceover"
	"net/http"
)

type DockerHubHooker struct {
	BaseHooker
}

func NewDockerHubHooker(baseHooker BaseHooker) *DockerHubHooker {
	return &DockerHubHooker{BaseHooker: baseHooker}
}

func (h *DockerHubHooker) Hook(w http.ResponseWriter, r *http.Request) {
	data, err := h.hookJsonBody(r)
	if err != nil {
		voiceover.Sayf("Docker hub hooker json decode failed: %s\n", err)
		_, _ = fmt.Fprintf(w, "fail")
		return
	} else {
		err = h.runCommand(data)
		if err != nil {
			voiceover.Sayf("Docker hub hooker run command failed: %s\n", err)
			_, _ = fmt.Fprintf(w, "fail")
			return
		} else {
			_, _ = fmt.Fprintf(w, "ok")
		}
	}}

func (h *DockerHubHooker) GetHttpMethods() routes.MethodCollection {
	return routes.MethodCollection{routes.POST}
}
//
//type dockerHubPayload struct {
//	CallbackURL string `json:"callback_url"`
//	PushData    struct {
//		Images   []string `json:"images"`
//		PushedAt float64  `json:"pushed_at"`
//		Pusher   string   `json:"pusher"`
//		Tag      string   `json:"tag"`
//	} `json:"push_data"`
//	Repository struct {
//		CommentCount    int64   `json:"comment_count"`
//		DateCreated     float64 `json:"date_created"`
//		Description     string  `json:"description"`
//		Dockerfile      string  `json:"dockerfile"`
//		FullDescription string  `json:"full_description"`
//		IsOfficial      bool    `json:"is_official"`
//		IsPrivate       bool    `json:"is_private"`
//		IsTrusted       bool    `json:"is_trusted"`
//		Name            string  `json:"name"`
//		Namespace       string  `json:"namespace"`
//		Owner           string  `json:"owner"`
//		RepoName        string  `json:"repo_name"`
//		RepoURL         string  `json:"repo_url"`
//		StarCount       int64   `json:"star_count"`
//		Status          string  `json:"status"`
//	} `json:"repository"`
//}
