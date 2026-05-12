package main

import "fmt"

type user struct {
	name string
	number int
}

type sender struct {
	user 
	rateLimit int
}

func test(s sender){
	fmt.Println(s.name)
	fmt.Println(s.number)
	fmt.Println(s.rateLimit)
	fmt.Println("=====================")
}

func main() {
	test(sender{
		user: user{
			name: "Mehul",
			number: 123,
		},
		rateLimit: 100,
	})

	test(sender{
		user: user{
			name: "Agarwal",
			number: 456,
		},
		rateLimit: 200,	
	})
}