package webserver

import (
    "fmt"
    "net/http"
)

type Webserver struct {
    router *router
    port string
    tokens chan string
}

func NewWebserver(p string) (*Webserver, error) {
    r := newrouter()
    w := Webserver{
            router: r,
            port: p,
            tokens: make(chan string)}
    return &w, nil
}

func (ws *Webserver) StartServer() {
    http.Handle("/", ws)

    fmt.Println("Starting server on port " + ws.port)
    go http.ListenAndServe(ws.port, nil)
    running := true
    for running {
        token := <-ws.tokens
        fmt.Println("Webserver got token " + token)
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
