package main

import (
	"html/template"
	"os"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

func main() {

	t := template.Must(template.New("template.html").ParseFiles("template.html"))

	err := t.Execute(os.Stdout, []Curso{{"Java", 20}, {"Go", 40}})

	if err != nil {
		panic(err)
	}
}
