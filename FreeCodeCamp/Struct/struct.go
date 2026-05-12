package main

import "fmt"

type car struct {

	brand string
	model string
	year int

	frontWheel wheel
	rearWheel wheel

}

type wheel struct{
	radius int
	alloy bool
}

func main(){

	myCar := car{
		brand: "Toyota",
		model: "Camry",
		year: 2020,
		frontWheel: wheel{
			radius: 18,
			alloy: true,
		},
		rearWheel: wheel{
			radius: 18,
			alloy: true,
		},
	}
	fmt.Println(myCar)

}