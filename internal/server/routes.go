package server

import (
	"encoding/json"
	"log"
	"net/http"

	"website/cmd/web"

	"github.com/a-h/templ"
)

func (s *Server) RegisterRoutes() http.Handler {

	r := http.NewServeMux()
	// r.HandleFunc("/", s.HelloWorldHandler)

	jsFiles := http.FileServer(http.FS(web.JSFiles))
	cssFiles := http.FileServer(http.FS(web.CSSFiles))
	r.Handle("/js/", jsFiles)
	r.Handle("/css/", cssFiles)
	r.Handle("/", templ.Handler(web.HelloForm()))
	// r.HandleFunc("/hello", web.HelloWebHandler)

	return r
}

func (s *Server) HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	resp := make(map[string]string)
	resp["message"] = "Hello World"

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("error handling JSON marshal. Err: %v", err)
	}

	_, _ = w.Write(jsonResp)
}
