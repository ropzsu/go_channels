

package main

import (
"os"
"reflect"
"strconv"
"fmt"
)


func main() {
  myPrint("Hello ", 42, "\n")	
  fmt.Printf("%#v\n", myPrint)
}

func myPrint(args ...interface{} ){
	for _, arg := range args {
	  switch v := reflect.ValueOf(arg); v.Kind() {
	    case reflect.String:
	       os.Stdout.WriteString(v.String())
	    case reflect.Int :
	       os.Stdout.WriteString(strconv.FormatInt(v.Int(), 10))

	  }
   }
} 
