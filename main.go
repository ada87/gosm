package main

import (
	"code.google.com/p/go.net/websocket"
	"fmt"
	"html/template"
	"net/http"
	"reflect"
	"strings"
)

type Control map[string]reflect.Value

var fns = make(Control)

var chttp = http.NewServeMux()

//webSocket总路由器
func socket(ws *websocket.Conn) {
	//	fmt.Println(os.FileMode)
	for {
		Rotesocket()
		var replay string
		if err := websocket.Message.Receive(ws, &replay); err != nil {
			fmt.Println("error")
		}
		if err := websocket.Message.Send(ws, "wocao"); err != nil {
			fmt.Println("error")
		}
	}

	//	fmt.Println("request")

}

func HandleIndex(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.ParseFiles("web/index.tpl")
	checkErr(err)
	tpl.Execute(w, nil)
}

//页面总路由
func Router(w http.ResponseWriter, r *http.Request) {
	//	fmt.Print(r.RequestURI)

	if strings.Contains(r.URL.Path, ".") {
		chttp.ServeHTTP(w, r)
		return
	}
	rote := strings.Split(r.RequestURI, "/")[1]
	if rote == "" {
		HandleIndex(w, r)
		return
	}
	fn, exists := fns[rote]
	if exists {
		fn.Call([]reflect.Value{reflect.ValueOf(w), reflect.ValueOf(r)})
	} else {
		fmt.Println("no router:", rote)
	}

}

func main() {
	sock := websocket.Handler(socket)
	http.HandleFunc("/", Router)
	http.Handle("/_", sock)
	http.Handle("/js/", http.FileServer(http.Dir("web")))
	http.Handle("/css/", http.FileServer(http.Dir("web")))
	http.Handle("/font/", http.FileServer(http.Dir("web")))
	for _, tab := range Config.Tabs {
		rote := tab.Handle
		methodName := strings.ToUpper(rote[:1]) + rote[1:]
		_, ok := reflect.TypeOf(&fns).MethodByName("Hander" + methodName)
		if ok {
			http.HandleFunc("/"+rote, Router)
			//			http.HandleFunc("/_"+rote, sock)
			fns[rote] = reflect.ValueOf(&fns).MethodByName("Hander" + methodName)
		} else {
			fmt.Println("no method :", methodName)
		}
	}
	if err := http.ListenAndServe(":9999", nil); err != nil {
		fmt.Printf("error")
	}
}

func init() {
	//	os.OpenFile("")
	//	fmt.Println(Config.Title)

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
