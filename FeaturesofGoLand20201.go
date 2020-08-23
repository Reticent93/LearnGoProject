package main

import (
	"fmt"
)

// Create the 'Rectangle' struct with 'width, height float64' fields
type Rectangle struct {
	width, height float64
}

// Create the 'Figure' interface with the 'Area() float64' method signature.
type Figure interface {
	Area() float64
}

//  The Area function counts area of the rectangle
func (r *Rectangle) Area() (string, float64) {
	text := "Rectangle area is: "
	//return 'text, r.width * r.height'
	return text, r.width * r.height
}

func main() {
	newRectangle := Rectangle{
		width:  5,
		height: 10,
	}
	fmt.Print(newRectangle.Area())
}
