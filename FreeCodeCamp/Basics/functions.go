package main

import "fmt"

func concat(s1 string, s2 string) string{
	return s1 + " " + s2;
}

//if  the type of the parameter are same then we can write like this , defint the datatype after the last variagble
func concat2(s1,s2 string) string{
	return s1 + " " + s2;
}

//Named return, we can name the return variable and then we can just return without mentioning the variable name
func concat3(s1,s2 string) (result string){
	result = s1 + " " + s2;
	return;
}

func main(){

}