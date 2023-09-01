package auth

import (
	"fmt"
	"log"
	"net/http"
)

type Server struct{}

func (s *Server) Start(port int, handler http.Handler) error {
	log.Printf("Starting server at http://127.0.0.1:%d", port)
	log.Print("Quit the server with CTRL-BREAK")

	return http.ListenAndServe(fmt.Sprintf(":%d", port), handler)
}
