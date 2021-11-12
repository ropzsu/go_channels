package main

import (
    "strings"
    "fmt"
)

func main() {
    // ZERO-VALUE:
    //
    // It's ready to use from the get-go.
    // You don't need to initialize it.
    var sb strings.Builder

    for i := 0; i < 1000; i++ {
        sb.WriteString("a")
    }

    fmt.Println(sb.String())
}

