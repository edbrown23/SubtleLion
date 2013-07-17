package main

import (
    "fmt"
    "webserver"
)

func main() {
    w, err := webserver.NewWebserver(8080)
    if err != nil {
        fmt.Println(err)
        panic("Idiot")
    }
    w.RegisterCallback("/test/", test)
    w.StartServer()
}

func test(i int) {
    fmt.Println(i)
}
