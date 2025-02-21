package middleware

import (
	"fmt"
	"net/http"
)

type Logger struct {
    handler http.Handler
}

func NewLogger(wrapedHandler http.Handler) *Logger {
    return &Logger{
        handler: wrapedHandler,
    }
}

func (l *Logger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
   l.handler.ServeHTTP(w, r)
   fmt.Printf("Requisição recebida %s %s", r.Method, r.URL.Path)
}
