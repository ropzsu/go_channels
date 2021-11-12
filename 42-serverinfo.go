package main

import (
        "fmt"
        "net/http"
        "os/exec"
        "strings"
        "time"
)

func serverInfo_ps(w http.ResponseWriter, req *http.Request) {
        psCMD := exec.Command("bash", "-c", "ps -e -o \"pid,etime,comm,args\" |grep -v ] |cat -n  ")
        psOut, _ := psCMD.Output()
        ts := time.Now().Format(time.RFC3339)
        fmt.Fprintf(w, fmt.Sprintf("[%s] msg \n%s", ts, psOut))
}

func serverInfo_disk(w http.ResponseWriter, req *http.Request) {
        psCMD := exec.Command("bash", "-c", " df -h |cat -n ")
        psOut, _ := psCMD.Output()

        ts := time.Now().Format(time.RFC3339)
        s := strings.ReplaceAll(string(psOut), "%", " disk rate")

        fmt.Fprintf(w, fmt.Sprintf("[%s] msg \n%s", ts, s))
}

func serverInfo_uptime(w http.ResponseWriter, req *http.Request) {
        psCMD := exec.Command("bash", "-c", " uptime ")
        psOut, _ := psCMD.Output()
        ts := time.Now().Format(time.RFC3339)
        fmt.Fprintf(w, fmt.Sprintf("[%s] msg \n%s", ts, psOut))
}

func main() {
        http.HandleFunc("/dnNvcHMK/ps", serverInfo_ps)
        http.HandleFunc("/dnNvcHMK/disk", serverInfo_disk)
        http.HandleFunc("/dnNvcHMK/uptime", serverInfo_uptime)

        http.ListenAndServe(":1998", nil)
}

