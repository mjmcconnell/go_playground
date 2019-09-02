package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
)

type Todo struct {
	Title string
	Done  bool
}

type TodoPageData struct {
	PageTitle string
	Todos     []Todo
}

type User struct {
	Name  string `json:name`
	Email string `json:email`
}

func logging(f http.HandlerFunc) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		log.Println(request.URL.Path)
		f(response, request)
	}
}

func home(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(response, "Hello, you've test requested: %s\n", request.URL.Path)
}

func test404(response http.ResponseWriter, request *http.Request) {
	urlVars := mux.Vars(request)
	fmt.Fprintf(response, "Test, %s", urlVars["path"])
}

func testTodo(response http.ResponseWriter, request *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/todo.html"))
	data := TodoPageData{
		PageTitle: "My TODO list",
		Todos: []Todo{
			{Title: "Task 1", Done: false},
			{Title: "Task 2", Done: true},
			{Title: "Task 3", Done: true},
		},
	}
	tmpl.Execute(response, data)
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

	router.HandleFunc("/", logging(home))
	router.HandleFunc("/json", logging(jsonView))
	testRouter.HandleFunc("/todo", logging(testTodo))
	testRouter.HandleFunc("/{path}", logging(test404))

	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(STATIC_DIR))))

	http.ListenAndServe(":8080", router)
}
