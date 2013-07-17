package webserver

import (
    "net/http"
    "reflect"
)

/*
Interacts with urls.go to route requests to their corresponding functions.
Wires the capture groups from a url's regex to the callbacks arguments
*/

type router struct {
    cbh *callbackHandler
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

func (r *router) routeRequest(req *http.Request) string {
    cb, err := r.cbh.findCallback(req.URL.Path)
    if err != nil {
        fmt.Println(err)
    }else {
        cbV := reflect.ValueOf(cb)
        args := []reflect.Value{reflect.ValueOf(7)}
        cbV.Call(args)
    }
}
