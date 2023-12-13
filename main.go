package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type Film struct {
	Title  string
	Director string
}

func main() {
	fmt.Println("Hello World!")
	films := map[string][]Film{
		"Films": {
			{Title: "The Godfather", Director: "Francis Ford Coppola"},
			{Title: "The Shawshank Redemption", Director: "Frank Darabont"},
		},
	}

	h1 := func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("index.html"))
		tmpl.Execute(w, films)
	}
	h2 := func(w http.ResponseWriter, r *http.Request) {
		log.Println("POSTED")
		title := r.PostFormValue("title")
		director := r.PostFormValue("director")
		tmpl := template.Must(template.ParseFiles("index.html"))
		tmpl.ExecuteTemplate(w, "film-list-element", Film{Title: title, Director: director})
	}
	http.HandleFunc("/", h1)
	http.HandleFunc("/add-film", h2)
	http.ListenAndServe(":8000", nil)
}
