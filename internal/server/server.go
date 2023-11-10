package server

import (
	"golang.org/x/net/websocket"
	"net/http"
	"websocket/internal/handler"
)

type Server struct {
	Port string
}

func New(port string) *Server {
	return &Server{
		Port: port,
	}
}

func (s *Server) Run(h *handler.Handler) {
	http.Handle("/ws", websocket.Handler(h.HandleWS))
	http.ListenAndServe(s.Port, nil)
}
