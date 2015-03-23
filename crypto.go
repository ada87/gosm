package main

import (
	"code.google.com/p/go.net/websocket"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"
	"io"
	"net/url"
	//	"fmt"
	"net/http"
	"reflect"
	"strings"
)

type Encryption struct {
	code string
	key  string
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

//aes 加密
func (e *Encryption) Encodeaes() string {
	src := []byte(e.code)
	if len(src)%aes.BlockSize != 0 {
		return "无法加密此字符串"
	}
	block, err := aes.NewCipher([]byte(e.key))
	if err != nil {
		return "密钥错误"
	}
	rtn := make([]byte, aes.BlockSize+len(src))
	iv := rtn[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "无法加密"
	}
	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(rtn[aes.BlockSize:], src)
	return hex.EncodeToString(rtn)
}

//aes 解密
func (e *Encryption) Decodeaes() string {
	decodeTxt, err := hex.DecodeString(e.code)
	if err != nil {
		return "解码失败"
	}
	if len(decodeTxt) < aes.BlockSize {
		return "无法解密此字符串"
	}

	block, err := aes.NewCipher([]byte(e.key))
	if err != nil {
		return "密钥错误"
	}
	iv := decodeTxt[:aes.BlockSize]
	decodeTxt = decodeTxt[aes.BlockSize:]

	if len(decodeTxt)%aes.BlockSize != 0 {
		return "无法解密此字符串"
	}
	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(decodeTxt, decodeTxt)
	return hex.EncodeToString(decodeTxt)
}

func (c *Controllers) PageCrypto(w http.ResponseWriter, r *http.Request) {
	Portal(w, "crypto.tpl", nil)
}
func (s *Sockets) SocketCrypto(ws *websocket.Conn, data map[string]string) {

	rtn := make(map[string]string)

	code, hascode := data["code"]
	if !hascode {
		rtn["result"] = "错误：请输入要加密的字符串"
		SocketReplay(ws, rtn)
		return
	}
	method, hasmethod := data["method"]
	if !hasmethod {
		method = "encode"
	}
	waygo, haswaygo := data["waygo"]
	if !haswaygo {
		waygo = "md5"
	}
	key, haskey := data["key"]
	if !haskey {
		waygo = "key"
	}

	var ec = Encryption{code, key}
	methodName := strings.ToUpper(method[:1]) + method[1:] + waygo
	var _, has = reflect.TypeOf(&ec).MethodByName(methodName)
	if has {
		fn := reflect.ValueOf(&ec).MethodByName(methodName)
		val := fn.Call([]reflect.Value{})
		rtn["result"] = val[0].String()
	} else {
		rtn["result"] = "错误：没有找到相应加密方法"
	}

	SocketReplay(ws, rtn)
}
