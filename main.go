package main

import (
	"fmt"
	"net/http"

	"github.com/vineethtw/vishu/handlers"
)

func main() {
	fmt.Printf("Starting application")
	http.Handle("/health", handlers.Health())
	http.ListenAndServe(":7887", nil)
}
