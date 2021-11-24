

package main

import (
"io"
"log"
"os"
"fmt"
"encoding/json"
)


type Lang struct {
	Name string
	Year int
	URL string
}

func do( f func(  Lang)) {
    input , err := os.Open("63-lang.json")
	if err != nil {
	   log.Fatal(err)
	}

	dec := json.NewDecoder(input)

	for {
	  var lang Lang
	  err := dec.Decode(&lang)
	  if err != nil {
	    if err == io.EOF {
	      break;
	    }
	    log.Fatal(err)
	  }
	  f(lang)
	}
}

func main() {
	do( func(lang Lang) {
	   fmt.Printf("%v\n", lang)
	})
}


