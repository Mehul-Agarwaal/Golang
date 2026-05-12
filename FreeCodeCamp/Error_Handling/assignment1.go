package main

import "fmt"

func sendSMSToCouple(msgTOCustomer, msgToSpouse string) (float64, error){
	costToCustomer, err := sendSMS(msgTOCustomer)
	if err != nil {
		return 0, fmt.Errorf("Failed to send SMS to customer: %v", err)
	}

	costToSpouse, err := sendSMS(msgToSpouse)
	if err != nil {
		return 0, fmt.Errorf("Failed to send SMS to spouse: %v", err)
	}

	return costToCustomer + costToSpouse, nil
}

func sendSMS(message string) (float64, error){
	const maxTextLen = 25
	const costPerChar = 0.002

	if(len(message) > maxTextLen){
		return 0, fmt.Errorf("Message is too long. Max length is %d characters", maxTextLen)
	}
	return float64(len(message)) * costPerChar, nil
}

func main(){
	msgToCustomer := "Hello, this is a message for you."
	msgToSpouse := "Hello, this is a message for your spouse."

	totalCost, err := sendSMSToCouple(msgToCustomer, msgToSpouse)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Printf("Total cost of sending SMS: $%.4f\n", totalCost)	
	
}
