package main

import (
	"code.google.com/p/go.net/websocket"
	"net/http"
)

func (c *Controllers) PageRegexp(w http.ResponseWriter, r *http.Request) {
	Portal(w, "regexp.tpl", Fields)
}

func (s *Sockets) SocketRegexp(ws *websocket.Conn, data map[string]string) {
	rst := execRegexp(data)
	SocketReplay(ws, rst)
}

func execRegexp(data map[string]string) map[string]string {
	var rtn = make(map[string]string)
	rtn["code"] = "0"
	act, _ := data["act"]
	switch act {
	case "new":
		fid, _ := data["fid"]
		fval, _ := data["fval"]
		fdes, _ := data["fdes"]
		err := insertVal(fid, fval, fdes)
		if err != nil {
			rtn["code"] = "1"
		}
	case "update":
		vid, _ := data["vid"]
		fval, _ := data["fval"]
		fdes, _ := data["fdes"]
		err := updateVal(vid, fval, fdes)
		if err != nil {
			rtn["code"] = "1"
		}
	}
	return rtn
}
