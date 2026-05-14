package main

import "fmt"

const(
	planFree = "Free"
	planPro = "Pro"
)

func getMessageWithRetriesForPlan(plan string)([]string, error){

	allMessages := getMessagesWithRetries()

	if plan == planFree{
		return allMessages[0:2],nil
	}else if plan == planPro{
		return allMessages,nil
	}
	return nil, fmt.Errorf("Invalid plan: %s", plan)

	
}

func getMessagesWithRetries() []string{
	return []string{"Message 1", "Message 2", "Message 3"}
}

func main(){

	messages, err := getMessageWithRetriesForPlan(planFree)
	if err != nil{
		fmt.Println("Error:", err)
	}else{
		fmt.Println("Messages for Free Plan:", messages)
	}

	messages, err = getMessageWithRetriesForPlan(planPro)
	if err != nil{
		fmt.Println("Error:", err)
	}else{
		fmt.Println("Messages for Pro Plan:", messages)
	}

	messages, err = getMessageWithRetriesForPlan("InvalidPlan")
	if err != nil{
		fmt.Println("Error:", err)
	}else{
		fmt.Println("Messages for Invalid Plan:", messages)
	}

}