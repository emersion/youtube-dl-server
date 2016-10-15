package server

import (
	"net/http"
	"os/exec"
)

type server struct {}

func New() http.Handler {
	return &server{}
}

func (s *server) handleInfo(w http.ResponseWriter, r *http.Request) {
	url := r.FormValue("url")

	cmd := exec.Command("youtube-dl", "--dump-json", "--all-subs", url)
	cmd.Stdout = w

	if err := cmd.Run(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.URL.Path {
	case "/info":
		s.handleInfo(w, r)
	default:
		http.NotFound(w, r)
	}
}
