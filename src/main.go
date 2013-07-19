package main

import (
    "fmt"
    "net/http"
    "webserver"
)

func main() {
    w, err := webserver.NewWebserver(8080)
    if err != nil {
        panic(err)
    }
    w.RegisterCallback("^/test/([a-zA-Z]+)$", stringTest)
    w.RegisterCallback("^/test/([\\d]+)$", numTest)
    w.RegisterCallback("^/test/([\\w]+)/([\\d]+)$", stringNumTest)
    w.StartServer()
}

func stringTest(r *http.Request, test string) {
    fmt.Println(test)
}

func numTest(r *http.Request, test int) {
    fmt.Println(test)
}

func stringNumTest(r *http.Request, test1 string, test2 int) {
    fmt.Printf("%s %d", test1, test2)
}
