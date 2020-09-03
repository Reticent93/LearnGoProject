package main

import (
	"html/template"
	"os"
)

var tpl *template.Template

func init() {

	tpl = template.Must(template.ParseGlob("templates/*"))

}

func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "two.gohtml", nil)
	if err != nil {
		panic(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "lasagne.php", nil)
	if err != nil {
		panic(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "anaheim.dland", nil)
	if err != nil {
		panic(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "lets.messup", nil)
	if err != nil {
		panic(err)
	}
}