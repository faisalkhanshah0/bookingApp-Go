// given 2 string s1 and s2 , we have to check whether they are anagram

// 2 strings are anagrm , if the chars are the same in both string

// freq of chars in the string are the same

//input Input: s1 = "danger" s2 = "garden"

package main

import (
	"fmt"
	"strings"
)

func main() {
	//recieving input for s1 & s2
	var s1 string
	var s2 string
	fmt.Println("Enter s1 value")
	fmt.Scan(&s1)
	fmt.Println("Enter s1 value")
	fmt.Scan(&s2)

	//calculating lengths of both string it should be equal in len
	if len(s1) == len(s2) {
		res := anagramfunc(s1, s2, len(s1))
		fmt.Println(res)
	} else {
		fmt.Println(false)
	}

	// fmt.Println(s1, s2)
}

func anagramfunc(s1 string, s2 string, count int) bool {
	arr := strings.Split(s1, "")
	fmt.Println(len(arr), cap(arr))
	counter := 0
	for _, v := range arr {
		// fmt.Println(i, v)
		if strings.Contains(s2, v) {
			counter++
		}
	}
	if counter == count {
		return true
	} else {
		return false
	}

}

// using map would be the solution

// {
// 	g : 1
// 	a : 1
// 	r : 1
// 	d : 1
// 	e : 1
// 	n : 1
// }
