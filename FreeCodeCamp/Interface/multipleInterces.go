package main

import (
	"fmt"
)

type expense interface {
	cost() float64
}

type printer interface {
	print()
}

type email struct {

	subscribed bool
	message string

}

func (e email) cost()  (totalCost float64) {
	if e.subscribed {
		totalCost = 0.1 * float64(len(e.message))
	}else{
		totalCost = 0.05 * float64(len(e.message))
	}
	return
}

func (e email) print() {
	fmt.Println(e.message)
}

func test(e expense, p printer) {
	fmt.Printf("Cost of the email: $%.2f\n", e.cost())
	fmt.Print("Email content: ")
	p.print()
}

func main() {
	email1 := email{subscribed: true, message: "Hello, this is a subscribed email."}
	email2 := email{subscribed: false, message: "Hello, this is an unsubscribed email."}

	test(email1, email1)
	test(email2, email2)
}