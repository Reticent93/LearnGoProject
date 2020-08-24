package main

import "fmt"

func main() {
	x := 42
	y := 43
	xi := []int{x, y} //slice of int

	a := "The morning sun"
	b := "The evening rocks"
	c := "Good morning Vietnam"
	xs := []string{a, b, c} //slice of string

	//Cannot combine the two
	fmt.Println(xi)
	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", y)
	fmt.Printf("%T\n", xi)

	fmt.Println(xs)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", c)
	fmt.Printf("%T\n", xs)
}
