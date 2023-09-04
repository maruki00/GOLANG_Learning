package main

import (
	"fmt"
	"net/http"
)

func handler(writer http.ResponseWriter, request *http.Request) {

	fmt.Fprintf(writer, "<html><body><b>Hello World, %s!</b></body></html>", request.URL.Path[1:])
	for p := range request.URL.Path {

		fmt.Fprintln(writer, p, "<br/>", request.URL.Path[p])
	}
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/get", handler)
	http.ListenAndServe(":8080", nil)
}
