
package main

import "fmt"

func say(name string) func ( int) string {
    var n = 0
	return func( x int ) string{
	   n++
	   return fmt.Sprintf("Hello: %s, job: %d Seq: %d " ,name,x, n )
	}

}

func main() {
	alex := say("brian")
	bob := say("jamie")

    for i := 0; i < 3; i ++ {
	  	fmt.Println(alex(101))
		fmt.Println(bob(102))  
    }

}

