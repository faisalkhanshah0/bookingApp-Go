package main

import (
	"fmt"
	"slices"
)

func main() {

	// {{[}]}
	// {[]{}}
	// {[]{[}]}
	// 12123210

	inputString := "{[]{}}()"

	setOfBraces := map[string]string{
		"(": ")",
		"{": "}",
		"[": "]",
		")": "(",
		"}": "{",
		"]": "[",
	}

	// arr1 = append(arr1, 1, 2, 3, 4, 5)
	// arr1 = slices.Delete(arr1, 1, 3)
	// arr1 = slices.Insert(arr1, 1, 2, 3)
	// fmt.Println(arr1)

	arr := []string{}
	for _, k := range inputString {
		ch := fmt.Sprintf("%c", k)

		if ch == "(" || ch == "{" || ch == "[" {
			fmt.Println("open : ", ch)
			arr = append(arr, ch)
		} else {
			if setOfBraces[ch] == arr[len(arr)-1] {
				fmt.Println("bf arr", arr, " got : ", ch, " cancelling", setOfBraces[ch], " arr len:", len(arr))
				arr = slices.Delete(arr, len(arr)-1, len(arr))
				fmt.Println("af arr", arr)
			}

		}
	}

	if len(arr) == 0 {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}

}
