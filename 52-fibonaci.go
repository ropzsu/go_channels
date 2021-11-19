
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


func run_job_1(r chan string) {
	t1 := time.Now()
	f := fibon()
	for i := 0 ; i < 10; i ++ {
	  fmt.Printf( "v1: fib(%d) = %d\n", i,  f())
	}
	t2 := time.Now()
	
	r <- "run1 done." + " cost time = " + fmt.Sprintf("cost time = %s\n", t2.Sub(t1))

}

func run_job_2( r chan string) {
	ch := make(chan int)
    t1  := time.Now()
	f2 := fibon2()
	for i := 0 ; i < 10; i ++ {
	  go f2(ch)
	  fmt.Printf( "v2: go routine: fib(%d) = %d\n", i,  <-ch )
	}
	t2  := time.Now()

	
	r <- "run2 done." + " cost time = " +fmt.Sprintf("cost time = %s\n", t2.Sub(t1))
}

func main() {

   r := make(chan string)
   go run_job_1(r)
   go run_job_2(r)
   fmt.Println("main: ", <- r)
   fmt.Println("main: ", <- r)
}



