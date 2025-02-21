package handler

import (
	"fmt"
	"io"
	"net/http"
)

func HelloHandler(w http.ResponseWriter, _r *http.Response) {
    fmt.Print("Hello!")
    io.WriteString(w, "Hello in page ihu")
}
