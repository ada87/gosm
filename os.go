package main

import (
	"code.google.com/p/go.net/websocket"
	"net"
	"net/http"
	"os"
	"os/exec"
	"os/user"
	"runtime"
	"strconv"
)

func (c *Controllers) PageOs(w http.ResponseWriter, r *http.Request) {
	Portal(w, "os.tpl", nil)
}

func (s *Sockets) SocketOs(ws *websocket.Conn, cmd []string) {
	rst := HandleResult(cmd)
	SocketReplay(ws, rst)
}

var allow = []string{"cmd", "winver", "write", "notepad", "calc", "mspaint", "mstsc", "devmgmt.msc", "services.msc", "taskmgr", "regedit", "compmgmt.msc", "fsmgmt.msc", "msconfig"}

func contain(str string, list []string) bool {
	for _, item := range list {
		if str == item {
			return true
		}
	}
	return false
}

//
//环境变量计算机情报
//网络信息
//用户信息
func HandleResult(cmd []string) map[string]string {
	var rtn = make(map[string]string)
	switch cmd[0] {
	case "env":
		envs := os.Environ()
		for key, item := range envs {
			rtn[strconv.Itoa(key)] = item
		}
	case "detail":
		//		memStat := new(runtime.MemStats)
		//		runtime.ReadMemStats(memStat)
		//		fmt.Println(memStat.Alloc)
		//		fmt.Println(runtime.CPUProfile())
		//		fmt.Println(runtime.Version())
		//		fmt.Println(runtime.NumCPU())

		rtn["CPU核心数"] = strconv.Itoa(runtime.NumCPU())
		//		cmd := exec.Command("cmd", "/c", "tasklist")
		//		buf, err := cmd.Output()
		//		cmd.Run()

		//		checkErr(err)
		//		fmt.Printf("%s", buf)

	case "net":
		adds, err := net.Interfaces()
		checkErr(err)
		for _, val := range adds {
			addrs, _ := val.Addrs()
			for _, addr := range addrs {
				rtn[addr.Network()] = addr.String()
			}
			addrs, _ = val.MulticastAddrs()
			for _, addr := range addrs {
				rtn[addr.Network()] = addr.String()
			}
			rtn["Mac"] = val.HardwareAddr.String()
			rtn["Name"] = val.Name
			rtn["Flag"] = val.Flags.String()
			//			rtn["Mtu"] = val.MTU
		}
	case "user":
		us, _ := user.Current()
		rtn["UserId"] = us.Uid
		rtn["GroupId"] = us.Gid
		rtn["Name"] = us.Name
		rtn["UserName"] = us.Username
		rtn["HomeDir"] = us.HomeDir
	default:
		if contain(cmd[0], allow) {
			exec.Command("cmd", "/c", cmd[0]).Run()
		} else {
			rtn["tips"] = "Usage: env, detail, net, user"
		}
	}
	return rtn
}
