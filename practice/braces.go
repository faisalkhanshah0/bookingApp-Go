package main

import (
	"fmt"
)

func main() {

	fmt.Println("welcom to bracecheck app")
	fmt.Println("Enter any bracket pattern like {[()]} ?")
	var rawstr string
	fmt.Scan(&rawstr)

	var pattern = map[string]string{
		"{": "}",
		"(": ")",
		"[": "]",
		"}": "{",
		")": "(",
		"]": "[",
	}
	// fmt.Println(pattern)

	if (len(rawstr) % 2) == 0 {
		fmt.Println("string is even. lets calculate")
		counter := 0
		for i := 0; i < (len(rawstr) / 2); i++ {
			c := fmt.Sprintf("%c", rawstr[i])
			v := pattern[c]
			j := len(rawstr) - (i + 1)
			m := fmt.Sprintf("%c", rawstr[j])
			// fmt.Printf("%v == %v", v, m)
			if v == m {
				counter++
			}

		}
		if counter == (len(rawstr) / 2) {
			fmt.Println("pattern is equal.")
		} else {
			fmt.Println("pattern is not valid.")
		}
	} else {
		fmt.Println("string is odd. exiting")
	}

}
