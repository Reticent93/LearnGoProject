package main

import "fmt"

func main() {
	x := "ABCDE"
	fmt.Println(x)

	fmt.Printf("%T", x)

	xr := []rune(x) //rune is int32
	fmt.Println(xr)

	for _, v := range xr {
		fmt.Printf("%d - %b - %#X\n", v, v, v) // this prints decimal, heximal, binary
	}

}
