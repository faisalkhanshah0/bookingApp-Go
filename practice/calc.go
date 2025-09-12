package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/PaesslerAG/gval"
)

// var globalvar int = 88

func main() {
	// fmt.Printf("value of global var is : %v \n", globalvar)
	add := func(a int, b int) int {
		return a + b
	}
	subtract := func(a int, b int) int {
		return a - b
	}
	multiply := func(a int, b int) int {
		return a * b
	}
	divide := func(a int, b int) (int, error) {
		if b == 0 {
			return 0, errors.New("cannot divide by zero.")
		}
		return a / b, nil
	}

	var eqn string
	var res int
	fmt.Println("Enter your equation below :")
	fmt.Scan(&eqn)
	trimmed := strings.TrimSpace(eqn)
	addeqn := strings.Split(trimmed, "+")
	subeqn := strings.Split(trimmed, "-")
	multiplyeqn := strings.Split(trimmed, "*")
	divideeqn := strings.Split(trimmed, "/")

	if len(addeqn) > 1 {
		fmt.Printf("%T", addeqn[0])
		a, _ := strconv.Atoi(addeqn[0])
		b, _ := strconv.Atoi(addeqn[1])
		res = add(a, b)
	} else if len(subeqn) > 1 {
		fmt.Println(subeqn, len(subeqn))
		a, _ := strconv.Atoi(subeqn[0])
		b, _ := strconv.Atoi(subeqn[1])
		res = subtract(a, b)
	} else if len(multiplyeqn) > 1 {
		fmt.Println(multiplyeqn, len(multiplyeqn))
		a, _ := strconv.Atoi(multiplyeqn[0])
		b, _ := strconv.Atoi(multiplyeqn[1])
		res = multiply(a, b)
	} else if len(divideeqn) > 1 {
		fmt.Println(divideeqn, len(divideeqn))
		a, _ := strconv.Atoi(divideeqn[0])
		b, _ := strconv.Atoi(divideeqn[1])
		res1, err := divide(a, b)
		if err != nil {
			fmt.Println("Error occured : ", err)
		}
		res = res1
	} else {
		fmt.Println("Invalid Input.")
	}
	fmt.Printf("Answer from my calc is : %v \n", strconv.Itoa(res))
	result, _ := gval.Evaluate(eqn, nil)
	fmt.Printf("Answer is : %v \n", result)
}
