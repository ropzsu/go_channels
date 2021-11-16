package main

import (
  "strings"
  "fmt"
  "os"
)

func main() {

  // "http://maps.google.com:9580,sa-saopaulo" 

  for i := 1 ; i < len(os.Args); i ++ {
       pair := strings.Split(os.Args[i], ",")
       fmt.Printf("%q => %q\n", pair[0], pair[1])
  }
}

