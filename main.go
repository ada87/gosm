package main

import (
	"code.google.com/p/go.net/websocket"
	"fmt"
	"net/http"
	"os"
)

//webSocket路由器
func socket(ws *websocket.Conn) {
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

//普通页面访问
func rote(w http.ResponseWriter, r *http.Request) {

}

func main() {
	http.HandleFunc("/", rote)
	sock := websocket.Handler(socket)
	http.Handle("/_ct", sock)
	if err := http.ListenAndServe(":9999", nil); err != nil {
		fmt.Printf("error")
	}
}

func init() {
	//	os.OpenFile("")

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
