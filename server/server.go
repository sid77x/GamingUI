package main
import ("net";"fmt";"os";"os/exec";"bufio";"time")
func main() {
	Conn()
}
func Conn(){
	var rurl string ="rtsp://localhost:5000\n"
	l, err := net.Listen("tcp", ":2568")
for{
	ErrChk(err)
	conn,err := l.Accept()
	ErrChk(err)
	r := bufio.NewReader(conn)
	line, _ := r.ReadString('\n')
	conn.Write([]byte(rurl))
	GameStr(line)
	Logger("Received:"+line)

	}
}
func Logger(info string){
	ti := time.Now().Format("2006-01-02 15:04:05")
	f,err := os.OpenFile("server_log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY , 0655)
	ErrChk(err)
	fmt.Fprintf(f,"%s",[]byte("["+ti+"]"+info+"\n"))
	f.Close()
}
func ErrChk(err error){
	if err != nil {
        fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
        os.Exit(1)
		}
}
func GameStr(gmn string){
	/*res, err := exec.Command("start scripts/"+gmn+".bat").CombinedOutput()
		fmt.Println("Running Game",gmn)
		ErrChk(err)
		fmt.Println(string(res))*/
	if gmn=="assc"{
		res, err := exec.Command("start scripts/"+gmn+".bat").CombinedOutput()
		fmt.Println("Running Game Assault Cube")
		ErrChk(err)
		fmt.Println(string(res))
	}
	if gmn=="pubg"{
		res, err := exec.Command("start scripts/"+gmn+".bat").CombinedOutput()
		fmt.Println("Running Game PUBG")
		ErrChk(err)
		fmt.Println(string(res))
	}
	if gmn=="fortnite"{
		res, err := exec.Command("start scripts/"+gmn+".bat").CombinedOutput()
		fmt.Println("Running Game Fortnite")
		ErrChk(err)
		fmt.Println(string(res))	
	}else{
		fmt.Println("Error")
	}
}