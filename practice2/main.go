package main

import (
	"fmt"
	"slices"
)

func main() {
	var arr = []int{}

	arr = append(arr, 1, 2, 3, 4, 5)
	arr = slices.Delete(arr, 1, 3)
	arr = slices.Insert(arr, 1, 2, 3)
	fmt.Println(arr)

	var arr1 = []int{3, 2, 4, 6, 4, 1}
	slices.Sort(arr1)
	arr1 = slices.Compact(arr1)
	fmt.Println(arr1)

	fmt.Println(slices.Compare(arr1, arr))

	arr3 := []any{"a", 2, 4, "b"}
	fmt.Println(arr3)

	//remove 1st element
	// arr4 := append([]any{}, arr3[1:]...)
	arr4 := arr3[1:]
	fmt.Println(arr4)

	//adding 1st element
	// arr4 = append([]any{"monkey"}, arr3...)
	arr4 = slices.Insert(arr3, 0, "monkey")
	fmt.Println(arr4...)

	//remove last element
	arr4 = append(arr4[:len(arr4)-1], 8, 9)
	fmt.Println(arr4)

	//insert at mid index
	arr4 = slices.Insert(arr4, 1, "b")
	fmt.Println(arr4)
	arr4 = slices.Delete(arr4, 1, 2)
	fmt.Println(arr4)
}
