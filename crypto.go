package main

import (
	"code.google.com/p/go.net/websocket"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"
	"net/url"
	//	"fmt"
	"net/http"
	"reflect"
	"strings"
)

type Encryption struct {
	code string
}

//MD5加密
func (e *Encryption) Encodemd5() string {
	buf := md5.Sum([]byte(e.code))
	pass := hex.EncodeToString(buf[:])
	return pass
}

//MD5解密
func (e *Encryption) Decodemd5() string {
	return "Md5目前无法解密，哈哈哈"
}

//Base64加密
func (e *Encryption) Encodebase64() string {
	buf := []byte(e.code)
	str := base64.StdEncoding.EncodeToString(buf)
	return str
}

//Base64解密
func (e *Encryption) Decodebase64() string {
	buf, err := base64.StdEncoding.DecodeString(e.code)
	if err != nil {
		return "解密失败"
	}
	return string(buf)
}

//url 打码
func (e *Encryption) Encodeurl() string {
	return url.QueryEscape(e.code)
}

//url 解码
func (e *Encryption) Decodeurl() string {
	rtn, err := url.QueryUnescape(e.code)
	if err != nil {
		return "解析URL失败"
	}
	return rtn
}

//sha1 加密
func (e *Encryption) Encodesha1() string {
	t := sha1.New()
	t.Write([]byte(e.code))
	return hex.EncodeToString(t.Sum(nil))
}

//sha1 解密
func (e *Encryption) Decodesha1() string {
	return "不好意思，无法破解"
}

//sha256 加密
func (e *Encryption) Encodesha256() string {
	t := sha256.New()
	t.Write([]byte(e.code))
	return hex.EncodeToString(t.Sum(nil))
}

//sha256 解密
func (e *Encryption) Decodesha256() string {
	return "不好意思，无法破解"
}

//sha512 加密
func (e *Encryption) Encodesha512() string {
	t := sha512.New()
	t.Write([]byte(e.code))
	return hex.EncodeToString(t.Sum(nil))
}

//sha512 解密
func (e *Encryption) Decodesha512() string {
	return "不好意思，无法破解"
}

func (c *Controllers) PageCrypto(w http.ResponseWriter, r *http.Request) {
	Portal(w, "crypto.tpl", nil)
}
func (s *Sockets) SocketCrypto(ws *websocket.Conn, cmd []string) {

	rtn := make(map[string]string)
	if len(cmd) < 3 {
		rtn["result"] = "错误：参数不够"
	} else {
		code := cmd[2]
		for i := 3; i < len(cmd); i++ {
			code += " " + cmd[i]
		}
		var ec = Encryption{code}
		methodName := strings.ToUpper(cmd[1][:1]) + cmd[1][1:] + cmd[0]
		var _, has = reflect.TypeOf(&ec).MethodByName(methodName)
		if has {
			fn := reflect.ValueOf(&ec).MethodByName(methodName)
			val := fn.Call([]reflect.Value{})
			rtn["result"] = val[0].String()
		} else {
			rtn["result"] = "错误：没有找到相应加密方法"
		}
	}
	SocketReplay(ws, rtn)
}
