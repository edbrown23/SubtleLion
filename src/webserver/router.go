package webserver

import (
    "fmt"
    "reflect"
)

/*
Interacts with urls.go to route requests to their corresponding functions.
Wires the capture groups from a url's regex to the callbacks arguments
*/

type router struct {
    *callbackHandler
}

func NewRouter() *router {
    cbh := newcallbackHandler()
    r := router{cbh}

    return &r
}

func (r *router) RegisterCallback(url string, callback interface{}) error {
    err := r.registerCallback(url, callback)
    if err != nil {
        return err
    }

    return nil
}

func (r *router) RouteRequest(url string) {
    cb, err := r.findCallback(url)
    if err != nil {
        fmt.Println(err)
    }else {
        cbV := reflect.ValueOf(cb)
        args := []reflect.Value{reflect.ValueOf(7)}
        cbV.Call(args)
    }
}
