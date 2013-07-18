package main

import (
    "fmt"
    "webserver"
)

func main() {
    w, err := webserver.NewWebserver(8080)
    if err != nil {
        panic(err)
    }
    w.RegisterCallback("^/test/(.+)$", test)
    w.StartServer()
}

func test(test string) {
    fmt.Println(test)
}
