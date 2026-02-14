package main

import (
	"html/template"
	"os"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

type Cursos []Curso

func main() {
	templates := []string{
		"../templates/header.html",
		"../templates/content.html",
		"../templates/footer.html",
	}

	t := template.Must(template.New("content.html").ParseFiles(templates...))
	err := t.Execute(os.Stdout, Cursos{
		{"Go", 40},
		{"Python", 30},
		{"JavaScript", 20},
	})
	if err != nil {
		panic(err)
	}
}
