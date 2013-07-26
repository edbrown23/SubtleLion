package webserver

import (
    "fmt"
    "net"
)

const (
    RECV_BUF_LEN = 1024
)

type ServerManager struct {
    ws *Webserver
    ctlPort string

}

func NewServerManager(cp string) *ServerManager {
    ws, err := NewWebserver("8080")
    if err != nil {
        panic(err)
    }

    s := ServerManager{ws: ws,
                       ctlPort: cp}
    return &s
}

func (sm *ServerManager) RunServer() {
    go sm.ws.StartServer()
    fmt.Println("Started game server on port 8080")
    sm.listen()
}

func (sm *ServerManager) listen() {
    ln, err := net.Listen("tcp", ":" + sm.ctlPort)
    if err != nil {
        fmt.Println("Could not listen on port " + sm.ctlPort)
        return
    }
    fmt.Println("Control Listening on port " + sm.ctlPort)
    running := true
    for running {
        conn, err := ln.Accept()
        if err != nil {
            fmt.Println("Error accepting connection")
            continue
        }
        running = sm.handleConnection(conn)
    }
}

func (sm *ServerManager) handleConnection(conn net.Conn) bool {
    fmt.Println("Parsing connection")
    buf := make([]byte, RECV_BUF_LEN)
    n, err := conn.Read(buf)
    if err != nil {
        conn.Write([]byte("No good"))
        return false
    }
    t := string(buf[:n])
    fmt.Println("Got access token " + t)
    sm.ws.tokens<- t
    return true
}
