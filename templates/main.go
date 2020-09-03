package main

import (
	"html/template"
	"log"
	"os"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.html"))
}

func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "main.go.html", nil)
	if err != nil {
		log.Fatal(err)
	}
	err = tpl.ExecuteTemplate(os.Stdout, "about.html", nil)
	if err != nil {
		log.Fatal(err)
	}

	f, err := os.Create("another.html")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
}
