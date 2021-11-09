
package main

import (
	"fmt"
	"io/ioutil"
	"os/exec"
	"time"
)

func main() {
   psCMD := exec.Command("bash" , "-c" , "ps -e -o \"pid,etime,comm,args\" ")
   psOut, _ := psCMD.Output()

   ts := time.Now().Format(time.RFC3339)  
   ioutil.WriteFile("/tmp/ps.log", append([]byte(ts+"\n"), psOut ... ), 0600)

   fmt.Printf("%s \n[%s] msg: %s\n",  
   	          string(psOut), ts, "Write  /tmp/ps.log ... done " )        
}


