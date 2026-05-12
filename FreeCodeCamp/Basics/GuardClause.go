package main

import (
	"fmt"
	"errors"
)

func divide (dividend, divisor int) (int , error){
	if divisor == 0 {
		return 0, errors.New("cannot divide by 0")

	}
	return dividend/divisor, nil
}

func main(){

	fmt.Println(divide(10,0));

	fmt.Println(divide(10,2));
}