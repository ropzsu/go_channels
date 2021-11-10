package main

import (
	"fmt"
	"math"
	"errors"
)

func Sqrt(x float64) (float64, error) {
    if (x < 0 ) {
	   return 0, errors.New(fmt.Sprintf("Error: cannot Sqrt negative number: %f", x))
	} 
	return math.Sqrt(x), nil
}

 
func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}

