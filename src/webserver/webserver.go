package webserver

import (
    "errors"
    "fmt"
    "net/http"
    "strconv"
)

type Webserver struct {
    router *router
    port int
}

func NewWebserver(p int) (*Webserver, error) {
    if p > 65535 || p < 0 {
        return nil, errors.New("Port must be within 0 - 65535")
    }

    r := newrouter()
    w := Webserver{
            router: r,
            port: p}
    return &w, nil
}

func (ws *Webserver) StartServer() {
    http.Handle("/", ws)
    ps := ":" + strconv.Itoa(ws.port)

    fmt.Println("Starting server on port " + ps)
    err := http.ListenAndServe(ps, nil)
    if err != nil {
        fmt.Println(err)
    }
}

func (ws *Webserver) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Got request from path: " + r.URL.Path)
    res := ws.router.routeRequest(r)
    fmt.Fprintf(w, res.Body)
}

func (ws *Webserver) RegisterAllCallbacks(cbs map[string]interface{}) error {
    for k, v := range cbs {
        err := ws.router.registerCallback(k, v)
        if err != nil {
            return err
        }
    }

    return nil
}

func (ws *Webserver) RegisterCallback(url string, callback interface{}) error {
    err := ws.router.registerCallback(url, callback)

    return err
}
