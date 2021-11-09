package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
     min := 0.0000001
	 
	 z := 1.0
	 for { 
	    fmt.Println(z)
	    if  math.Abs( z*z - x ) < min {
		  return z
		} else {
		   z -= (z*z - x )/ (2*x)
		}
	 }
}

func main() {
	fmt.Println(Sqrt(2))
}

