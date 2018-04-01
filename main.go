package main

import (
	"fmt"
	"net/http"
)

func healthHandler() http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "OK")
	})
}

func main() {
	fmt.Printf("Starting application")
	http.Handle("/health", healthHandler())
	http.ListenAndServe(":7887", nil)
}
