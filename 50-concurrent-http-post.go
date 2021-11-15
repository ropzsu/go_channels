package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func http_post(addr string, post_message string) (r string, err error) {
	t1 := time.Now()
	postBody := bytes.NewBuffer([]byte(post_message))

	http_client := http.Client{
		Timeout: 2 * time.Second,
	}

	resp, err := http_client.Post(addr, "application/json", postBody)
	if err != nil {
		t2 := time.Now()
		return fmt.Sprintf("[%s] Post  %s [%s] msg: %s cost time: %s", t2.Format(time.RFC3339), addr, "failed.", "Post request  failed.", t2.Sub(t1)), err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t2 := time.Now()
		return fmt.Sprintf("[%s] Post %s [%s] msg: %s cost time: %s", t2.Format(time.RFC3339), addr, "failed.", "Read body failed.", t2.Sub(t1)), err
	}

	defer resp.Body.Close()
	t2 := time.Now()
	return fmt.Sprintf("[%s] Post %s [%s] msg: %s  cost time: %s", t2.Format(time.RFC3339), addr, "success", body, t2.Sub(t1)), nil
}

func main() {

	post_request := ` {"email":"Toby@example.com","name":"Toby"}  `

	if len(os.Args) < 2 {
		return
	} else {
		ch := make(chan string, 1)

		for i := 1; i < len(os.Args); i++ {
			url := os.Args[i]
			go func(u string) {
				r, _ := http_post(u, post_request)
				ch <- r
			}(url)
		}

		//check result
		for i := 1; i < len(os.Args); i++ {
			fmt.Println(<-ch)
		}
	}
}
