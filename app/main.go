package main

import (
    "fmt"
    "net/http"
    "github.com/gorilla/mux"
)

func main() {
    router := mux.NewRouter()

    router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello, you've test requested: %s\n", r.URL.Path)
    })

    router.HandleFunc("/test", func(response http.ResponseWriter, request *http.Request) {
        fmt.Fprintf(response, "Test, %s", request.URL.Path)
    })
    router.HandleFunc("/test/{foo}", func(response http.ResponseWriter, request *http.Request) {
        urlVars := mux.Vars(request)

        fmt.Fprintf(response, "Test, %s", urlVars["foo"])
    })

    http.ListenAndServe(":8080", router)
}
