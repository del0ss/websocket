package main

import (
	"websocket/internal/handler"
	"websocket/internal/server"
)

const (
	port = ":8080"
)

func main() {
	s := server.New(port)
	h := handler.New()
	s.Run(h)
}
