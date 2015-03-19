package main

import (
	"code.google.com/p/go.net/websocket"
	"fmt"
	"net/http"
	"os"
)

func (c *Controllers) PageOs(w http.ResponseWriter, r *http.Request) {
	Portal(w, "os.tpl", nil)
}

func (s *Sockets) SocketOs(ws *websocket.Conn) {
	os.Environ()

	websocket.Message.Send(ws, "haha")
}
