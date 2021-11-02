package main

import "fmt"
import "time"

func check(u string, ch chan<- bool) {
    time.Sleep(4 * time.Second)
    ch <- true
}

func IsReachable(urls []string) bool {
    ch := make(chan bool, len(urls))
    for _, url := range urls {
        go check(url, ch)
    }
    time.AfterFunc(time.Second, func() { ch <- false })
    return <-ch
}
func main() {
    fmt.Println(IsReachable([]string{"url1", "url2"}))
}

