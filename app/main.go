package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	testRouter := router.PathPrefix("/test").Subrouter()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, you've test requested: %s\n", r.URL.Path)
	})

	testRouter.HandleFunc("", func(response http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(response, "Test, %s", request.URL.Path)
	})
	testRouter.HandleFunc("/{foo}", func(response http.ResponseWriter, request *http.Request) {
		urlVars := mux.Vars(request)

		fmt.Fprintf(response, "Test, %s", urlVars["foo"])
	})

	http.ListenAndServe(":8080", router)
}
