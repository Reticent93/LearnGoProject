package main

import "fmt"

func main() {
	xi := []int{1, 2, 4, 5, 7, 8, 9}
	//foo(xi)//this is [1,2,4,5,7,8,9]
	foo(xi...) //spread operator
}

func foo(i ...int) {
	//func foo(ii []int)
	//fmt.Println(ii)//this is [1,2,4,5,7,8,9]
	fmt.Println(i)
	fmt.Printf("%T\n", i)
}
