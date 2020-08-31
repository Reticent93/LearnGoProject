package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	//tpl, err := template.ParseFiles("tpl.gohtml")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//err = tpl.Execute(os.Stdout, nil)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//Do not use the above code in production

	tpl, err := template.ParseFiles("one.txt", "two.txt") //this reads the files
	if err != nil {
		log.Fatal("uh-oh ", err)

		tpl.ExecuteTemplate(os.Stdout, "one.txt", nil)
		tpl.ExecuteTemplate(os.Stdout, "two.txt", "Greg")

	}
}
