package main

import "fmt"

type person struct {
	first  string
	last   string
	saying string
}

//func (receiver) identifier(params) (returns) {code}
func (p person) speak() { //attach the value speak to any person
	fmt.Println(p.first, "says", p.saying)
}

func (sa secretAgent) speak() { //attach the value speak to any person
	fmt.Println(sa.first, "says wackabacka haha", sa.saying)
}

type secretAgent struct {
	person
	ltk bool
}

type human interface { //defines functionality
	speak() //anything that implements this func automatically implements this interface
}

func foo(h human) {
	h.speak()
}

func main() {
	p1 := person{
		first:  "James",
		last:   "Bond",
		saying: "Shaken, not stirred",
	}
	p2 := person{
		first:  "Jenny",
		last:   "MoneyPenny",
		saying: "I laugh in the face of danger",
	}

	sa1 := secretAgent{
		person: person{
			first:  "Ian",
			last:   "Fleming",
			saying: "The books were based on true stories",
		},
		ltk: true,
	}

	//different way to say the same things
	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println("<------------>")

	fmt.Println(p1.first, "says", p1.saying)
	fmt.Println(p1.first, "says", p2.saying)

	fmt.Println("<------------>")

	p1.speak()
	p2.speak()

	fmt.Println("<-----------> interface")

	foo(p1)
	foo(p2)
	foo(sa1)
}
