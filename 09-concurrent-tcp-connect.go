
package main

import (
	"fmt"
	"net"
	"os"
	"time"
)


func connect_only(addr string)  (resp string, err error) {
    t1 := time.Now()
    d := net.Dialer{Timeout: time.Second *2 }
	conn, err := d.Dial("tcp", addr)
	t2 := time.Now()
	if err != nil {
		return fmt.Sprintf( "[%s] msg: %s %s %s cost time: %s", t2.Format(time.RFC3339),  "connect ", addr , "[failed]", t2.Sub(t1)), err
	}
	defer conn.Close()	
	return fmt.Sprintf( "[%s] msg: %s %s %s cost time: %s", t2.Format(time.RFC3339), "connect ", addr , "[success]", t2.Sub(t1)), nil
}

func main() {
   if (len(os.Args) < 2 ) {
      return 
   } else {
      ch := make(chan string, 1)

      for i := 1; i < len(os.Args) ; i ++ {
        url := os.Args[i]        
       go func() {
         r, _ := connect_only(url)
         ch <-  fmt.Sprintf("%s", r)               
       }()
      }

      //check result 
      for i := 1; i < len(os.Args) ; i ++ {
         fmt.Println(<-ch) 
      }
   }
}



