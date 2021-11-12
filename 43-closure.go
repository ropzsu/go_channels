
package main

import "fmt"

func say(name string) func () string {
	return func() string{
	   return fmt.Sprintf("Hello: %s" ,name)
	}

}

func main() {
	alex := say("alex")
	bob := say("bob")

	fmt.Println(alex())
	fmt.Println(bob())
}


