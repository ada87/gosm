package main

import (
	"code.google.com/p/go.net/websocket"
	"crypto/md5"
	"encoding/hex"
	//	"fmt"
	"net/http"
	"reflect"
	"strings"
)

type Encryption struct {
	code string
}

func (e *Encryption) Encodemd5() string {
	buf := md5.Sum([]byte(e.code))
	pass := hex.EncodeToString(buf[:])
	return pass
}
func (e *Encryption) Decodemd5() string {
	return "Md5目前无法解密，哈哈哈"
}
func (c *Controllers) PageCrypto(w http.ResponseWriter, r *http.Request) {
	Portal(w, "crypto.tpl", nil)
}
func (s *Sockets) SocketCrypto(ws *websocket.Conn, cmd []string) {

	rtn := make(map[string]string)
	if len(cmd) < 3 {
		rtn["错误"] = "参数不够"
	} else {
		var ec = Encryption{cmd[2]}
		methodName := strings.ToUpper(cmd[1][:1]) + cmd[1][1:] + cmd[0]
		var _, has = reflect.TypeOf(&ec).MethodByName(methodName)
		if has {
			fn := reflect.ValueOf(&ec).MethodByName(methodName)
			val := fn.Call([]reflect.Value{})
			rtn["结果"] = val[0].String()
		} else {
			rtn["错误"] = "没有找到相应加密方法"
		}
	}
	SocketReplay(ws, rtn)
}
