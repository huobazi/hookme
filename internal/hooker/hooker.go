package hooker

import (
	"bufio"
	"bytes"
	"encoding/json"
	"github.com/huobazi/hookme/pkg/routes"
	"github.com/huobazi/hookme/pkg/voiceover"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"syscall"
)

type Hooker interface {
	GetRequestPath() string
	GetHttpMethods() routes.MethodCollection
	Hook(w http.ResponseWriter, r *http.Request)
}

type BaseHooker struct {
	Name        string
	RequestPath string
	WorkDir     string
	Command     string
}

func (h BaseHooker) GetRequestPath() string {
	return h.RequestPath
}

func (h BaseHooker) GetHttpMethods() routes.MethodCollection {
	return routes.MethodCollection{routes.POST}
}

func (h *BaseHooker) runCommand(args interface{}) (err error) {
	cmdPath, err := exec.LookPath(h.Command)
	if err != nil {
		cmdPath, err = exec.LookPath(filepath.Join(h.WorkDir, h.Command))
	}
	if err != nil {
		voiceover.Sayf("Error locating command: '%s'", h.Command)
		return err

	}

	cmd := exec.Command(cmdPath)
	cmd.Dir = h.WorkDir

	stdin, err := cmd.StdinPipe()
	if err != nil {
		return err
	}
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}
	defer stdout.Close()
	stderr, err := cmd.StderrPipe()
	if err != nil {
		return err
	}
	defer stderr.Close()

	go readIo(stdout)
	go readIo(stderr)

	go func() {
		defer stdin.Close()
		payload, _ := json.Marshal(args)
		b := bytes.NewBuffer(payload)
		_, writeError := b.WriteTo(stdin)
		if pathError, ok := writeError.(*os.PathError); ok && pathError.Err == syscall.EPIPE {
		} else if writeError != nil {
			voiceover.Sayf("Exec command failed: %s\n", writeError)
		}
	}()

	err = cmd.Run()
	if err != nil {
		return err
	}

	return nil

}

func readIo(r io.Reader) {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		voiceover.Say(scanner.Text())
	}
}