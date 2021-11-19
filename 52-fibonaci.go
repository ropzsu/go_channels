
package main 

import (
"fmt"
"time"
)
func fibon() func () int {
	var i, j = 0, 1

	return func () int {
	    v := i
	    i, j = j , i+j
	    return v

	}
}

func main() {
    t1 := time.Now()
	f := fibon()
	for i := 0 ; i < 10; i ++ {
	  fmt.Printf( "fib(%d) = %d\n", i,  f())
	}
	t2 := time.Now()

	fmt.Printf("cost time = %s\n", t2.Sub(t1))

}

/*
[brianguo@VM-16-7-centos go_channels]$ go run 11-fibon.go
fib(0) = 0
fib(1) = 1
fib(2) = 1
fib(3) = 2
fib(4) = 3
fib(5) = 5
fib(6) = 8
fib(7) = 13
fib(8) = 21
fib(9) = 34
cost time = 33.937µs

*/
