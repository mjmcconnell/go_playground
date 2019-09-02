package main

import (
	"fmt"
	"html/template"
	"github.com/gorilla/mux"
	"net/http"
)

type Todo struct {
	Title string
	Done bool
}

type TodoPageData struct {
    PageTitle string
    Todos     []Todo
}



func main() {
	STATIC_DIR := "static"

	router := mux.NewRouter()
	testRouter := router.PathPrefix("/test").Subrouter()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, you've test requested: %s\n", r.URL.Path)
	})

	testRouter.HandleFunc("", func(response http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(response, "Test, %s", request.URL.Path)
	})

	testRouter.HandleFunc("/todo", func(response http.ResponseWriter, request *http.Request) {
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

	})
	router.HandleFunc("/{path}", func(response http.ResponseWriter, request *http.Request) {
		urlVars := mux.Vars(request)
		fmt.Fprintf(response, "404, could not find path: %s", urlVars["path"])
	})

	// fs := http.FileServer(http.Dir("static/"))
 	//    router.Handle("/static/", http.StripPrefix("/static/", fs))

 	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(STATIC_DIR))))

	http.ListenAndServe(":8080", router)
}
