package main

import "fmt"

func main() {
	foo()
	bar()
}
func foo() {
	for i := 0; i <= 100; i++ {
		fmt.Println(i)
	}
}

func bar() {
	for i := 0; i <= 100; i++ {
		if i%2 == 0 { // conditional that only prints even numbers
			fmt.Println(i)

		}
	}
}
