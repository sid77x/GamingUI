package main

import "C"

import ("net";"fmt";"bufio";"os/exec";"os";"time";"strings";"github.com/zserge/webview")
func main() {}
func Logger(info *C.char){
	ti := time.Now().Format("2006-01-02 15:04:05")
	f,err := os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY , 0655)
	ErrChk(err)
	fmt.Fprintf(f,"%s",[]byte("["+ti+"]"+C.GoString(info)+"\n"))
	f.Close()
}
//export Connect
func Connect(gac *C.char){
	var err error
	gmcode := bufio.NewReader(strings.NewReader(C.GoString(gac)))
	conn, err := net.Dial("tcp", ":2568")
	ErrChk(err)
	response := bufio.NewReader(conn)
	send, err := gmcode.ReadString('\n')
	Logger(C.CString(send))
	ErrChk(err)
	conn.Write([]byte(send))
	res, err := response.ReadString('\n')
	ErrChk(err)
	GameCl(C.CString(res))
}
func ErrChk(err error){
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s \n", err.Error())
		os.Exit(1)
		}
}
//export GameCl
func GameCl(url *C.char){
	res, err := exec.Command("start ga/gacli.exe "+C.GoString(url)+"client.rel.conf").CombinedOutput()
	ErrChk(err)
	defer fmt.Println(string(res))
}
//export UI
func UI(title *C.char, url *C.char) {
	webview.Open(C.GoString(title),
		C.GoString(url), 850, 80, true)
}