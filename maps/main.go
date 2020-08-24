package main

import "fmt"

func main() {
	m := map[string]int{"James": 007, "Jenny": 8} //maps a string that returns an int
	fmt.Printf("%T\n", m)
	fmt.Println(m)

	fmt.Println(m["James"])
	fmt.Println(m["Jenny"])
}
