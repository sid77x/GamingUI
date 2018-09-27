package main

import ("github.com/zserge/webview";"os";"fmt";"os/exec")
func main() {
	serv,err := exec.Command("python","web.py").CombinedOutput()
	ErrChk(err)
	defer fmt.Println(serv)
	webview.Open("Cloud Gaming",
		"http://localhost:5000", 850, 80, true)
}
func ErrChk(err error){
	if err != nil {
        fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
        os.Exit(1)
		}
}