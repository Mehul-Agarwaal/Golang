package main

import "fmt"

type authenticationInfo struct {
	username string
	password string
}

func (a authenticationInfo)  getBasicAuth() string {
	return fmt.Sprintf("AUTHENTICATION: %s:%s", a.username, a.password)
}

func test (auth authenticationInfo) {
	fmt.Println(auth.getBasicAuth())
}

func main () {

	authInfo := authenticationInfo {
		username: "admin",
		password: "admin123",
	}	

	test(authInfo)
}
	