package main

import (
	"fmt"
)

func main() {
	var number int
	fmt.Println("Enter an number to print the table.")
	fmt.Scan(&number)
	for i := 1; i <= 10; i++ {
		fmt.Printf("%v X %v = %v \n", number, i, number*i)
	}
}
