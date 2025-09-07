package main

import (
	"fmt"
	"time"
)

func Login(name string, pass string) bool {
	time.Sleep(2 * time.Second)
	isLoggedin := false

	if name == "shah" && pass == "root" {
		isLoggedin = true
		fmt.Println("Login successfull")
		fmt.Println("Welcome back ", name)

	} else {
		fmt.Println("Login failed. try again")
		isLoggedin = false
	}
	return isLoggedin

}
