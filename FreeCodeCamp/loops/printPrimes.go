package main

import "fmt"

func printPrimes(n int){
	for i:=2;i<=n;i++{
		isPrime := true
		for j:=2;j<i;j++{
			if i%j == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			fmt.Printf("%d ", i)
		}
	}
}

func main(){
	n := 20
	fmt.Printf("Prime numbers up to %d: ", n)
	printPrimes(n)
}