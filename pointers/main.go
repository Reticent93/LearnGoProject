package main

import "fmt"

func main() {
	x := 42
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	fmt.Println(&x) //& shows the address of the value

	y := &x
	fmt.Println(y)
	fmt.Printf("%T\n", y) //points to the address
	fmt.Println(*y)       //Gives the value of what y is pointing to. In this case it's x

	*y = 43        // This is changing the value that y is pointing at
	fmt.Println(x) // value changes from 42 to 43
}
