package main

import (
	"fmt"
	"net/http"
)

func index(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "index handler, %s!", request.URL.Path[1:])
}

func error(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "error")
}
