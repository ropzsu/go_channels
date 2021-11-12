
package main

import (
        "fmt"
        "net/http"
        "os/exec"
        "time"
)

func make_handle_func(cmd string) func (  http.ResponseWriter,   *http.Request) {
        return func  (w http.ResponseWriter, req *http.Request) {
                psCMD := exec.Command("bash", "-c",  cmd )
                psOut, _ := psCMD.Output()
                ts := time.Now().Format(time.RFC3339)
                fmt.Fprintln(w, fmt.Sprintf("[%s] msg \n", ts) , string(psOut) )
        }
}

func main() {
        http.HandleFunc("/dnNvcHMK/ps", make_handle_func("ps -e -o  \"pid,etime,comm,args\" |cat -n "))
        http.HandleFunc("/dnNvcHMK/disk", make_handle_func("df -h |cat -n ")  )
        http.HandleFunc("/dnNvcHMK/uptime", make_handle_func("uptime") )
        http.ListenAndServe(":1998", nil)
}

