package handler

import (
	"fmt"
	"io"
	"net/http"
)

func helloHandler(w http.ResponseWriter) {
    fmt.Print("Hello!")
    io.WriteString(w, "Hello in page ihu")
}
