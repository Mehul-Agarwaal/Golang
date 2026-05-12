package main

import (
	"fmt"
	"math"
)	

type expense interface {
	cost()  (totalCost float64)
}

type email struct {
	to string
	isSubscription bool
	body string
}

type sms struct {
	toPhoneNumber string
	isSubscription bool
	body string
}

func (e email) cost() (totalCost float64) {
	if e.isSubscription {
		return totalCost = 0.0
	}	
	return 1.0
}

func (s sms) cost() (totalCost float64) {
	if s.isSubscription {
		return totalCost = 5.0
	}
	return 10.0
}

func getExpenseReport(e expense) (string, float64) {
	em, ok := e.(email)
	if ok {
		return fmt.Sprintf("Email to %s", em.to), math.Round(em.cost()*100)/100
	}

	s, ok := e.(sms)
	if ok {
		return fmt.Sprintf("SMS to %s", s.toPhoneNumber), math.Round(s.cost()*100)/100
	}

	return "Unknown expense", 0.0
}

/* 
func getExpenseReport(e expense) (string, float64) {
	switch v := e.(type) {
	case email:
		return fmt.Sprintf("Email to %s", v.to), math.Round(v.cost()*100)/100
	case sms:
		return fmt.Sprintf("SMS to %s", v.toPhoneNumber), math.Round(v.cost()*100)/100
	}
	default:
		return "Unknown expense", 0.0
}

func main() {

	email1 := email{
		to: "<EMAIL>",
		isSubscription: false,
		body: "Hello, this is a test email.",
	}

	sms1 := sms{
		toPhoneNumber: "+1234567890",
		isSubscription: true,
		body: "Hello, this is a test SMS.",
	}

	fmt.Println(getExpenseReport(email1))
	fmt.Println(getExpenseReport(sms1))

}