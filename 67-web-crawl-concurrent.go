

package main

import (
"io"
"log"
"os"
"fmt"
"encoding/json"
"net/http"
"io/ioutil"
"time"
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

func count(name, url string) {
	start := time.Now()
	r, err := http.Get(url)
	if err != nil {
	  fmt.Printf("%s:%s \n", name, err)
	  return 
	}
	n, _ := io.Copy(ioutil.Discard, r.Body)
	r.Body.Close()
	fmt.Printf("%s %d [%.2fs]\n", name, n, time.Since(start).Seconds())
}


func main() {
	
	do( func(lang Lang) {
      go count(lang.Name, lang.URL)
	})
	time.Sleep(10 * time.Second)
	
}




