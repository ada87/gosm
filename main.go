package main

import (
	"bytes"
	"code.google.com/p/go.net/websocket"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"os/exec"
	"reflect"
	"strconv"
	"strings"
)

//控制器集合
type Controllers map[string]reflect.Value

//Socket处理器集合
type Sockets map[string]reflect.Value

var cs = make(Controllers)
var ss = make(Sockets)

//webSocket总路由器
func socket(ws *websocket.Conn) {
	var err error
	for {
		var recive []byte
		var data map[string]string
		if err = websocket.Message.Receive(ws, &recive); err != nil {
			break
		}
		json.Unmarshal(recive, &data)
		rote, _ := data["path"]
		fn, exists := ss[rote]
		if exists {
			fn.Call([]reflect.Value{reflect.ValueOf(ws), reflect.ValueOf(data)})
		} else {
			fmt.Println("no socket: ", rote)
		}
		if err != nil {
			break
		}
	}
}

//静态文件处理器
var chttp = http.NewServeMux()

//页面总路由
func Router(w http.ResponseWriter, r *http.Request) {
	if strings.Contains(r.URL.Path, ".") {
		chttp.ServeHTTP(w, r)
		return
	}
	rote := strings.Split(r.RequestURI, "/")[1]
	if rote == "" {
		HandleIndex(w, r)
		return
	}
	fn, exists := cs[rote]
	if exists {
		fn.Call([]reflect.Value{reflect.ValueOf(w), reflect.ValueOf(r)})
	} else {
		w.WriteHeader(404)
		fmt.Println("no router:", rote)
	}
}

//程序入口，绑定 页面总路由器  和 websocket总路由器
func main() {
	sock := websocket.Handler(socket)
	http.HandleFunc("/", Router)
	http.Handle("/_", sock)
	http.Handle("/js/", http.FileServer(http.Dir("web")))
	http.Handle("/css/", http.FileServer(http.Dir("web")))
	http.Handle("/css-source/", http.FileServer(http.Dir("web")))
	http.Handle("/images/", http.FileServer(http.Dir("web")))
	http.Handle("/font/", http.FileServer(http.Dir("web")))

	for _, tab := range Config.Tabs {
		rote := tab.Handle
		methodName := strings.ToUpper(rote[:1]) + rote[1:]
		//Page Handder
		var _, ok = reflect.TypeOf(&cs).MethodByName("Page" + methodName)
		if ok {
			http.HandleFunc("/"+rote, Router)
			cs[rote] = reflect.ValueOf(&cs).MethodByName("Page" + methodName)
		}

		//Socket Handder

		_, ok = reflect.TypeOf(&ss).MethodByName("Socket" + methodName)
		if ok {
			//			http.HandleFunc("/"+rote, Router)
			ss[rote] = reflect.ValueOf(&ss).MethodByName("Socket" + methodName)
		}

	}
	//TODO 一个。此句话只在window系统下有效。
	var err = exec.Command("cmd", "/c", "start", "http://localhost:"+strconv.Itoa(Config.Port)+"/").Start()
	checkErr(err)
	err = http.ListenAndServe(":"+strconv.Itoa(Config.Port), nil)
	checkErr(err)
}

type PageData struct {
	Config config
	Data   interface{}
}

func (page PageData) Partial(tpl string) string {
	//	fmt.Println(tpl)
	t, err := template.ParseFiles("web/" + tpl)
	checkErr(err)
	var buf bytes.Buffer
	t.Execute(&buf, page)
	//	return t.ExecuteTemplate()
	return buf.String()
}

//模板引擎
func Portal(w http.ResponseWriter, tpl string, data interface{}) {
	pd := PageData{Config, data}
	w.WriteHeader(200)
	t, err := template.ParseFiles("web/"+tpl, "web/content/head.tpl", "web/content/foot.tpl")
	checkErr(err)
	t.Execute(w, pd)
}

//webSocket引擎
func SocketReplay(ws *websocket.Conn, data interface{}) {
	replay, err := json.Marshal(data)
	checkErr(err)
	websocket.Message.Send(ws, string(replay))
}

//进入首页
func HandleIndex(w http.ResponseWriter, r *http.Request) {
	Portal(w, "index.tpl", nil)
}

//异常处理，暂时只做panic
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
