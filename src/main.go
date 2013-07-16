package main

import (
    "fmt"
    "webserver"
)

func main() {
    w := webserver.NewRouter()
    w.RegisterCallback("hello", test)
    w.RouteRequest("hello")
}

func test(i int) {
    fmt.Println(i)
}
