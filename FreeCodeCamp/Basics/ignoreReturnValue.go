package main

import "fmt"

func getName() (string,string){
	return "Mehul","Agarwal";
}

func main(){

	firstName, lastName := getName(); // we can ignore the second return value by using _ (underscore)
	fmt.Println(firstName, lastName)

	firstName2, _ := getName(); // we can ignore the second return value by using _ (underscore)
	fmt.Println(firstName2)

	_, lastName2 := getName(); // we can ignore the first return value by using _ (underscore)
	fmt.Println(lastName2)
}