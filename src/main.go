package main

import (
    "webserver"
)

func main() {
    w := webserver.NewUrls()
    w.RegisterCallback("hello", test)
    
}

func test() {

}