package hooker

import (
	"fmt"
	"github.com/huobazi/hookme/pkg/routes"
	"github.com/huobazi/hookme/pkg/voiceover"
	"net/http"
)

type AliyunRegistryHooker struct {
	BaseHooker
}

func NewAliyunRegistryHooker(baseHooker BaseHooker) *AliyunRegistryHooker {
	return &AliyunRegistryHooker{BaseHooker: baseHooker}
}

//
//func (h *AliyunRegistryHooker) Hook(w http.ResponseWriter, r *http.Request) {
//	defer r.Body.Close()
//
//	decoder := json.NewDecoder(r.Body)
//	var data aliyunRegistryPayload
//	err := decoder.Decode(&data)
//	if err != nil {
//		voiceover.Sayf("Aliyun registry hooker json decode failed: %s\n", err)
//		_, _ = fmt.Fprintf(w, "fail")
//		return
//	} else {
//		err = h.runCommand(data)
//		if err != nil {
//			voiceover.Sayf("Aliyun registry hooker run command failed: %s\n", err)
//			_, _ = fmt.Fprintf(w, "fail")
//			return
//		} else {
//			_, _ = fmt.Fprintf(w, "ok")
//		}
//	}
//}

func (h *AliyunRegistryHooker) Hook(w http.ResponseWriter, r *http.Request) {
	data, err := h.hookJsonBody(r)
	if err != nil {
		voiceover.Sayf("Aliyun registry hooker json decode failed: %s\n", err)
		_, _ = fmt.Fprintf(w, "fail")
		return
	} else {
		err = h.runCommand(data)
		if err != nil {
			voiceover.Sayf("Aliyun registry hooker run command failed: %s\n", err)
			_, _ = fmt.Fprintf(w, "fail")
			return
		} else {
			_, _ = fmt.Fprintf(w, "ok")
		}
	}
}

func (h *AliyunRegistryHooker) GetHttpMethods() routes.MethodCollection {
	return routes.MethodCollection{routes.POST}
}

//
//type (
//	aliyunRegistryPayload struct {
//		PushData   struct {
//			Digest   string `json:"digest"`
//			PushedAt string `json:"pushed_at"`
//			Tag      string `json:"tag"`
//		}   `json:"push_data"`
//		Repository struct {
//			DateCreated            string `json:"date_created"`
//			Name                   string `json:"name"`
//			Namespace              string `json:"namespace"`
//			Region                 string `json:"region"`
//			RepoAuthenticationType string `json:"repo_authentication_type"`
//			RepoFullName           string `json:"repo_full_name"`
//			RepoOriginType         string `json:"repo_origin_type"`
//			RepoType               string `json:"repo_type"`
//		} `json:"repository"`
//	}
//)
