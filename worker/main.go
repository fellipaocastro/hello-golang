package main

import (
    "fmt"
    "time"

    "github.com/fellipecastro/hello-golang/message"
)

func main() {
    for {
        fmt.Println(message.Hello)
        time.Sleep(10 * time.Second)
    }
}
