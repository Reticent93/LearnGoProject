package main

import "fmt"

func main() {
	defer foo() //will run last
	bar()
}
func foo() {
	fmt.Println("foooooo")

}

func bar() {
	fmt.Println("barrrrr") //will run first
	defer one()            //will run third
	two()                  //will run second
}

func one() {
	fmt.Println("oneeee")
}
func two() {
	fmt.Println("twoooo")
}
