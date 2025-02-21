package main

import (
	"assets/handler"
	"fmt"
	"net/http"
)

type ServerHttp struct {
    port string // ex: ":9000"
    muxServer *http.ServeMux
}

func NewHttpServer(port string, muxServer *http.ServeMux) *ServerHttp {
    return &ServerHttp{
        port: port,
        muxServer: muxServer,
    }
}

func (s *ServerHttp) run() {
    http.ListenAndServe(s.port, s.muxServer)
    fmt.Printf("Server rodando na porta: %s", s.port)
}

func main() {
    mux1 := http.NewServeMux()

    /* Handlers */
    mux1.Handle("/", handler.HelloHandler)

    server := NewHttpServer(":8080", mux1)
    server.run()
}
