

package main

import (
"io"
"log"
"os"
)


type Lang struct {
	Name string
	Year int
	URL string
}

func main() {
    input , err := os.Open("lang.json")
	if err != nil {
	   log.Fatal(err)
	}
	io.Copy(os.Stdout, input)
}

