
package main

import "fmt"

func say(name string) func () string {
    var n = 0
	return func() string{
	   n++
	   return fmt.Sprintf("Hello: %s, Seq: %d" ,name, n)
	}

}

func main() {
	alex := say("brian")
	bob := say("jamie")

    for i := 0; i < 3; i ++ {
	  	fmt.Println(alex())
		fmt.Println(bob())  
    }

}

