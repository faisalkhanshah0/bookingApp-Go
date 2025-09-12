package main

import (
	"flag"
	"fmt"
	"slices"
	"strings"
)

func main() {
	strref := flag.String("string", "naman", "palindrome check for string")
	flag.Parse()
	str := *strref
	fmt.Println("Enterred string is : ", str)
	// var str string
	var rev string
	// fmt.Scan(&str)
	strarr := strings.Split(str, "")
	slices.Reverse(strarr)
	rev = strings.Join(strarr, "")
	if str == rev {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}
