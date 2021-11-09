
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

   ioutil.WriteFile("/tmp/ps.log", psOut, 0600)

   fmt.Printf("%s \n[%s] msg: %s\n",  string(psOut), 
              time.Now().Format(time.RFC3339),  
              "Write  /tmp/ps.log ... done " )        
}


