package main

import "fmt"

func main() {
	y := foo()
	fmt.Println(y)
	b := func(z string) {
		fmt.Println(z)
	}
	b("james")
	b("jenny")

}

func foo() string {
	return "this is foo"
}
