package webserver

import (
    "fmt"
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

func Newrouter() *router {
    cbh := newcallbackHandler()
    r := router{cbh}

    return &r
}

func (r *router) registerCallback(url string, callback interface{}) error {
    err := r.cbh.registerCallback(url, callback)

    return err
}

func (r *router) routeRequest(req *http.Request) string {
    cb, subs, err := r.cbh.findCallback(req.URL.Path)
    if err != nil {
        return "404 - No handler for path: " + req.URL.Path
    }else {
        fmt.Println(subs)
        cbV := reflect.ValueOf(cb)
        // The first arg is the string itself if there's a match
        args := convertToReflectValues(subs[1:])
        cbV.Call(args)
    }

    return "OK"
}

func convertToReflectValues(args []string) []reflect.Value {
    o := make([]reflect.Value, len(args))
    for i := range args {
        o[i] = reflect.ValueOf(args[i])
    }
    return o
}
