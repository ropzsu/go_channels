package main

import (
	"fmt"
	"errors"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func http_get(addr string) (r string, err error) {
	t1 := time.Now()
	resp, err := http.Get(addr)
	if err != nil {
		t2 := time.Now()
		return fmt.Sprintf("[%s] Get %s [%s] msg: %s cost time: %s", t2.Format(time.RFC3339), addr, "failed.", "connet failed.", t2.Sub(t1)),
			errors.New("connet failed.")

	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t2 := time.Now()
		return fmt.Sprintf("[%s] Get %s [%s] msg: %s cost time: %s", t2.Format(time.RFC3339), addr, "failed.", "Read body failed.", t2.Sub(t1)),
			errors.New("Read Body failed.")
	}

	defer resp.Body.Close()
	t2 := time.Now()
	return fmt.Sprintf("[%s] Get %s [%s] msg: %s  cost time: %s", t2.Format(time.RFC3339), addr, "success", body, t2.Sub(t1)), err
}

func main() {
	if len(os.Args) < 2 {
		return
	} else {
		ch := make(chan string, 1)

		for i := 1; i < len(os.Args); i++ {
			url := os.Args[i]
			go func(u string) {
				r, _ := http_get(u)
				ch <- r
			}(url)
		}

		//check result
		for i := 1; i < len(os.Args); i++ {
			fmt.Println(<-ch)
		}
	}
}

