
package main 

import "fmt"

func normal_ret(s string) string {
	return s
}

func go_channel_ret(s string , ch chan string) {
	ch <- s 
}


func main () {
	ch := make(chan string , 1)
	r1 := normal_ret("hello world")
	go go_channel_ret("hello, go channel", ch)

	fmt.Printf("r1 = \"%s\" \n", r1)
	fmt.Printf("r2 = \"%s\" \n", <-ch)

}

