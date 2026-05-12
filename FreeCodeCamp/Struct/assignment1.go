package main

import "fmt"

type msgToSend struct {
	msg string
	to user
	from user
}

type user struct {
	name string
	email string
}

func canSendMessage(m msgToSend) bool {
	if m.to.email == "" || m.from.email == "" || m.msg == ""|| m.to.name == "" || m.from.name == "" {
		return false
	}
	return true;
}

func main(){
	u1 := user {
		name: "Mehul",
		email: "abc",
	}
	u2 := user {
		name: "Agarwal",
		email: "xyz",
	}
	u3:= user {
		name: "John",
		email: ""	,
	}
	m1 := msgToSend{
		msg: "Hello",
		to: u1,
		from: u2,
	}
	m2 := msgToSend{
		msg: "Hello",
		to: u1,
		from: u3,
	}
	fmt.Println(canSendMessage(m1))

	fmt.Println(canSendMessage(m2))	

}