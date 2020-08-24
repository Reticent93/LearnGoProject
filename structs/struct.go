package main

import "fmt"

type hotdog int

var x hotdog

type person struct {
	first   string
	license int
	sayings []string
}

func main() {

	p1 := person{
		first:   "James",
		license: 007,
		sayings: []string{"Shaken, not stirred", "Bond, James Bond"},
	}
	x = 7
	fmt.Printf("%T\n", x)
	fmt.Println(x)

	y := int(x)
	fmt.Printf("%T\n", y)
	fmt.Println(y)

	fmt.Println(p1)
}
