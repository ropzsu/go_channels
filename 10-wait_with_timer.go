package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"
)

func working(name string, c chan string) {
	t1 := time.Now()
	time.Sleep(time.Millisecond * 1300)
	c <- "[success] msg: job -  " + name + " done , cost time: " + fmt.Sprintf("%s", time.Now().Sub(t1))
}

func working_with_timeout(d time.Duration) (resp string, err error) {
	ch := make(chan string, 1)

	go func() { working("hello world", ch) }()

	select {
	case r := <-ch:
		return r, nil
	case <-time.After(d):
		return "[failed] msg: timeout " + fmt.Sprintf("%s", d), errors.New("Timeout")
	}

}

func main() {
	timeout := 1500
	if len(os.Args) > 1 {
		timeout, _ = strconv.Atoi(os.Args[1])
	}
	d := time.Millisecond * time.Duration(timeout)

	resp, _ := working_with_timeout(d)
	t := time.Now().Format(time.RFC3339)
	fmt.Printf("[%s] %s\n", t, resp)
}
