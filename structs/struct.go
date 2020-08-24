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

	p2 := person{
		first:   "Jenny",
		license: 8,
		sayings: []string{"When Bond can't, I can", "I laugh in the face of danger"},
	}
	fmt.Println(p2)

	xp := []person{p1, p2}
	fmt.Println(xp)

	mp := map[string]person{"Mr": p1, "Ms": p2}
	fmt.Println(mp)

	x = 7
	fmt.Printf("%T\n", x)
	fmt.Println(x)

	y := int(x)
	fmt.Printf("%T\n", y)
	fmt.Println(y)

	fmt.Println(p1)
}
