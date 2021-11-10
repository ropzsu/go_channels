package main

import (
    "fmt"
    "net/http"
    "os/exec"
    "time"
)

func serverInfo(w http.ResponseWriter, req *http.Request) {
   psCMD := exec.Command("bash" , "-c" , "ps -e -o \"pid,etime,comm,args\" |grep -v ] |cat -n  ")
   psOut, _ := psCMD.Output()

   upCMD := exec.Command("bash" , "-c" , "  uptime ")
   upOut, _ := upCMD.Output()

   ts := time.Now().Format(time.RFC3339)  
   w.Header().Set("Content-Type", "text/html; charset=utf-8")

  refresh := `<script>
    window.setInterval('refresh()', 1000); 
    function refresh() {   window .location.reload();  }
    </script>
   `;

   fmt.Fprintf(w, fmt.Sprintf("%s <body><pre>[%s] msg \n%s \n[%s] msg: %s</pre></body>\n", refresh,  ts, psOut, ts , upOut ) )
}


func main() {
    http.HandleFunc("/", serverInfo)
    http.ListenAndServe(":1998", nil)
}

