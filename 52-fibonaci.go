
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

func fibon2() func (chan int)   {
	var i, j = 0, 1

	return func ( c chan int)   {
	    v := i
	    i, j = j , i+j
	    c <- v
	}
}


func main() {
    t1 := time.Now()
	f := fibon()
	for i := 0 ; i < 10; i ++ {
	  fmt.Printf( "v1: fib(%d) = %d\n", i,  f())
	}
	t2 := time.Now()

	fmt.Printf("cost time = %s\n", t2.Sub(t1))


    ch := make(chan int)
    t1  = time.Now()
	f2 := fibon2()
	for i := 0 ; i < 10; i ++ {
	  go f2(ch)
	  fmt.Printf( "v2: go routine: fib(%d) = %d\n", i,  <-ch )
	}
	t2  = time.Now()

	fmt.Printf("cost time = %s\n", t2.Sub(t1))

}
