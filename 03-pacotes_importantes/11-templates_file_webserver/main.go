package main

import (
	"html/template"
	"net/http"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t := template.Must(template.New("template.html").ParseFiles("template.html"))

		err := t.Execute(w, []Curso{{"Java", 20}, {"Go", 40}})

		if err != nil {
			panic(err)
		}
	})

	http.ListenAndServe(":8000", nil)
}
