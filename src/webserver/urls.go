package webserver

import (
    "fmt"
)

var urls = map[string]interface{} {
    "/hello/": Test,
}

func HandleUrl(url string) {
    urls[url].(func())()
}



func Test() {
    fmt.Println("Testies")
}