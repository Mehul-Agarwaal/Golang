package main

import "fmt"

type rect struct {
	length int
	width int
}

func (r rect) area() int {
	return r.length * r.width
}

func main() {

	r:= rect {
		length: 10,
		width: 5,
	}

	fmt.Println("Area of rectangle is: ", r.area())

}