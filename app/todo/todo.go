package todo

import (
	"html/template"
	"net/http"
)

type todo struct {
	Title string
	Done  bool
}

type todoPageData struct {
	PageTitle string
	Todos     []todo
}

func TodoView(response http.ResponseWriter, request *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/todo.html"))
	data := todoPageData{
		PageTitle: "My TODO list",
		Todos: []todo{
			{Title: "Task 1", Done: false},
			{Title: "Task 2", Done: true},
			{Title: "Task 3", Done: true},
		},
	}
	tmpl.Execute(response, data)
}
