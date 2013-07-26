package main

import (
    "fmt"
    "net/http"
    "webserver"
)

func main() {
    // w, err := webserver.NewWebserver("8080")
    // if err != nil {
    //     panic(err)
    // }
    // w.RegisterCallback("^/test/([a-zA-Z]+)$", stringTest)
    // w.RegisterCallback("^/test/([\\d]+)$", numTest)
    // w.RegisterCallback("^/test/([\\w]+)/([\\d]+)$", stringNumTest)
    // w.StartServer()
    s := webserver.NewServerManager("9500")
    s.RunServer()
}

func stringTest(r *http.Request, test string) webserver.HttpResponse {
    return webserver.HttpResponse{test, 200}
}

func numTest(r *http.Request, test int) webserver.HttpResponse {
    body := fmt.Sprint(test)
    return webserver.HttpResponse{body, 200}
}

func stringNumTest(r *http.Request, test1 string, test2 int) webserver.HttpResponse {
    body := fmt.Sprintf("%s %d", test1, test2)
    return webserver.HttpResponse{body, 200}
}
