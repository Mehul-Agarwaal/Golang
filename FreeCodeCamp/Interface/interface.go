package main

import (
	"fmt"
	"time"
)

type message interface {
	getMessage() string
	
}

type birthdayMessage struct {
	reciever string
	birthdayTime time.Time
}

type sendingReport struct {
	reportName string
	numberOfSends int
}

func (b birthdayMessage) getMessage() string {
	return fmt.Sprintf("Happy Birthday %s! Your birthday is on %s", b.reciever, b.birthdayTime.Format("January 2, 2006"))
}

func (s sendingReport) getMessage() string {
	return fmt.Sprintf("Report: %s has been sent %d times", s.reportName, s.numberOfSends)
}

func test(m message) {
	fmt.Println(m.getMessage())
}

// This code defines an interface called 'message' which has a method 'getMessage()' that returns a string. 
// Two structs, 'birthdayMessage' and 'sendingReport', implement this interface by providing their own versions of the 'getMessage()' method. 
// The 'test' function takes an argument of type 'message' and prints the result of calling 'getMessage()' on it. 
// In the main function, you can create instances of 'birthdayMessage' and 'sendingReport', and pass them to the
func main(){

	test(birthdayMessage{
		reciever: "Alice",
		birthdayTime: time.Date(1990, time.May, 15, 0, 0, 0, 0, time.UTC),
	})
	test(sendingReport{
		reportName: "Monthly Sales",
		numberOfSends: 5,
	})


}
