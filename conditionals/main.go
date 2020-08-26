package main

import "fmt"

func main() {
	ee := []int{2, 3, 4, 7, 9}
	fmt.Println(ee)

	for i, v := range ee {
		fmt.Println(i, v)
		if i == 3 {
			fmt.Println("HEY, i is EQUAL TO", i)
		}
	}
	n := "Q"
	switch n {
	case "Moneypenny":
		fmt.Println("miss money")
	case "Bond":
		fmt.Println("bond james")
	case "Q":
		fmt.Println("this is q")
	default:
		fmt.Println("This is default")
	}

	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(!true)

}
