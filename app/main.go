package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/mjmcconnell/go_playground/logging"
	"github.com/mjmcconnell/go_playground/todo"
	"net/http"
)

type User struct {
	Name  string `json:name`
	Email string `json:email`
}

func home(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(response, "Hello, you've test requested: %s\n", request.URL.Path)
}

func test404(response http.ResponseWriter, request *http.Request) {
	urlVars := mux.Vars(request)
	fmt.Fprintf(response, "Test, %s", urlVars["path"])
}

func jsonView(response http.ResponseWriter, request *http.Request) {
	bob := User{
		Name:  "bob",
		Email: "bob@example.com",
	}
	json.NewEncoder(response).Encode(bob)

}

func main() {
	STATIC_DIR := "static"

	router := mux.NewRouter()
	testRouter := router.PathPrefix("/test").Subrouter()

	router.HandleFunc("/", logging.Logger(home))
	router.HandleFunc("/json", logging.Logger(jsonView))
	router.HandleFunc("/todo", logging.Logger(todo.TodoView))
	testRouter.HandleFunc("/{path}", logging.Logger(test404))

	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(STATIC_DIR))))

	http.ListenAndServe(":8080", router)
}
