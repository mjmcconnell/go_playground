package logging

import (
	"log"
	"net/http"
)

func Logger(f http.HandlerFunc) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		log.Println(request.URL.Path)
		f(response, request)
	}
}
