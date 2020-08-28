package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {

	//assign to variable
	name := "Greg"

	//CREATE STRING
	//string print
	str := `html here ` + name + ` more html`
	fmt.Println(str)

	s := fmt.Sprint(`mas ` + name + ` menos`)
	fmt.Println(s)

	//CREATE FILE
	//io.Copy to file
	nf, err := os.Create("newfile.txt")
	if err != nil {
		log.Fatal("Whoops", err)

	}
	io.Copy(nf, strings.NewReader(s))

	//CREATE FILE
	//writestring to file
	nf2, err := os.Create("newfile.txt")
	if err != nil {
		log.Fatal("whhooops2", err)
	}

	n, err := nf2.WriteString(str)
	if err != nil {
		log.Fatal("Whoops again", err)
	}
	fmt.Println("bytes written", n)

}
