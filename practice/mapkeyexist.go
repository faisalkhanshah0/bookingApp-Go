package main

import "fmt"

func main() {
	arr := []string{"a", "b", "c", "d", "e"}
	for i, v := range arr {
		fmt.Println(i, v)
	}

	mapvar := map[string]string{
		"a": "animals",
		"b": "birds",
		"c": "fishes",
	}
	var key string
	fmt.Println("enter your key you want ot check :")
	fmt.Scan(&key)

	//first way

	if mapvar[key] != "" {
		fmt.Println("key exist")
	} else {

		fmt.Println("not exist")
	}

	//second way
	val, ok := mapvar[key]
	if ok {
		fmt.Println("Alice exists with value:", val)
	} else {
		fmt.Println("Alice does not exist")
	}

}

//pprof : package to check memorey leakages and performance of go routines.
