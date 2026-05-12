package main

import "fmt"

func main(){
	//while loop

	i:=1
	for i<=3{
		fmt.Println(i)
		i = i + 1
	}

	//infinite loop
	for {
		fmt.Println("Infinite loop")
	}
}