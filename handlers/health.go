package handlers

import (
	"fmt"
	"net/http"
)

/*
Health provides a health check endpoint
*/
func Health() http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "OK")
	})
}
