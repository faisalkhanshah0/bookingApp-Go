package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {

	// 1st approach
	dummytext := os.Args[1]
	// dummytext := "Lorem		Ipsum		is sim\nply dummy 92 "
	// dummytext := "Lorem		Ipsum	is sim\nply dummy 92 text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum."
	re := regexp.MustCompile(`\w+`)
	// re := regexp.MustCompile(`\S+`)
	arr := re.FindAllString(dummytext, -1)
	fmt.Println(arr, "word count in given string is : ", len(arr))

	// 2nd approach
	re1 := regexp.MustCompile(`\s+`)
	text := strings.TrimSpace(dummytext)
	arr1 := re1.Split(text, -1)
	fmt.Println(arr1, "word count in given string is : ", len(arr1))

	// 3rd approach
	arr2 := strings.Fields(dummytext)
	fmt.Println("wc in given string is : ", len(arr2))

	// 4th approach

	// arr := strings.Split(dummytext, " ")
	// wc := len(arr)
	// counter := 0
	// for _, v := range arr {
	// 	if strings.Contains(v, "	") {
	// 		occurance := strings.Count(v, "	")
	// 		counter = counter - occurance
	// 	}
	// 	if v == "" {
	// 		counter++
	// 	}
	// }
	// wc = wc - counter

	// fmt.Printf("word count of given text is %v \n", wc)
}
