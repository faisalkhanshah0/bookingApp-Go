package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	numraw := os.Args[1]
	num, _ := strconv.Atoi(numraw)
	if (num % 2) == 0 {
		fmt.Printf("given no %v is even \n", num)
	} else {
		fmt.Printf("given no %v is odd \n", num)
	}

	var factraw int
	fmt.Println("enter any number to calculate factorial ?")
	fmt.Scan(&factraw)
	res := fact(factraw)
	fmt.Printf("factorial of given no %v is %v \n", factraw, res)
}

func fact(factraw int) int {
	var fac int = 1
	for i := factraw; i > 0; i-- {
		fac = fac * i
	}
	return fac
}
