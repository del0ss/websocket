package handler

import (
	"fmt"
	"golang.org/x/net/websocket"
	"io"
	"log"
)

type Handler struct {
	Conns map[*websocket.Conn]bool
}

func New() *Handler {
	return &Handler{
		Conns: make(map[*websocket.Conn]bool),
	}
}

func (h *Handler) HandleWS(ws *websocket.Conn) {
	fmt.Println("New connection: ", ws.RemoteAddr())

	h.Conns[ws] = true
	h.ReadLoop(ws)
}

func (h *Handler) ReadLoop(ws *websocket.Conn) {
	buf := make([]byte, 1024)

	for {
		n, err := ws.Read(buf)

		if err != nil {
			if err != io.EOF {
				break
			}
			log.Print("Ошибка прочтения!")
			return
		}

		msg := buf[:n]
		fmt.Println(string(msg))

		ws.Write([]byte("Thx you for the message!!!!"))
	}
}
