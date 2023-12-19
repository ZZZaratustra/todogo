package main

import (
	"net/http"
	"text/template"
)

var tmpl *template.Template

type Todo struct {
  Item string
  Done bool
}

type PageData struct {
  Title string
  Todos []Todo
}

func todo (w http.ResponseWriter, r *http.Request) {
  data:= PageData{
    Title: "TODO LIST",
    Todos: []Todo{
       {Item:"Install go", Done:true},
      {Item:"Install go", Done:true},
    },
    
  }
}

func main() {
  mux := http.NewServeMux()
  tmpl = template.Must(template.ParseFiles("templates/index.gohtml"))
  mux.HandleFunc("/todo", todo)
}
