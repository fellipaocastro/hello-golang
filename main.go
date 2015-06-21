package main

import (
    "fmt"
    "flag"
    "log"
    "net/http"

    "github.com/fellipecastro/hello-golang/message"
)

func main() {
    port := flag.String("port", "8080", "HTTP Port")
    flag.Parse()

    log.Println("Starting Server on", *port)
    http.HandleFunc("/", hello)
    log.Fatal(http.ListenAndServe(":"+*port, nil))
}

func hello(rw http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(rw, message.Hello)
}
