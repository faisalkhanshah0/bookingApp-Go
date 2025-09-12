package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"slices"
	"strings"
)

var globalvar = "global"

func main() {

	fmt.Println("############ Variables #############")

	var localvar = "localmain"
	var localvar1 string = "localvar1main"
	localvar2 := "localvar2"
	age := 30
	isGradudated := true
	const pi = 3.14
	vartochange := 10
	fmt.Println("value before :", vartochange)
	callbyreference(&vartochange)
	fmt.Println("value after :", vartochange)

	fmt.Println(globalvar)
	fmt.Println(localvar, localvar1, localvar2, age, isGradudated, pi)
	callbyvalue(localvar2)

	fmt.Println("############ Arrays #############")

	blankarr := [5]string{} //blank array without assignment
	blankarr[0] = "a"
	blankarr[1] = "b"
	blankarr[2] = "c"
	blankarr[3] = "d"
	blankarr[4] = "e"
	arr := [5]string{"apple", "banana", "guava", "pinapple", "grapes"}
	fmt.Println(blankarr, arr, arr[0])
	// blankarr[4] = "" // unset any element
	copy(arr[0:], arr[1:]) //unseting 1st element unshift
	arr[len(arr)-1] = ""   // remove last elemtn part of unshift

	fmt.Println("After all : ", blankarr, arr, arr[0])
	fmt.Println(arr[:2])
	fmt.Println(arr[3:])
	fmt.Println(arr[:])

	fmt.Println("############ Slice #############")

	blankslice := []string{}
	blankslice = append(blankslice, "a")
	fmt.Println(blankslice)
	slice := []string{"dog", "cat", "camel", "elephant", "lion", "tiger"}
	slice = append(slice, "rabbit") // adding element at last
	fmt.Println(slice[:3])
	fmt.Println(slice)
	slice = append([]string{"monkey"}, slice...) // adding element at first
	fmt.Println(slice)
	fmt.Printf("some slice operation : %T \n", slice)
	slice = slice[1:] // remove 1st element
	fmt.Println(slice)
	slice = slice[0 : len(slice)-1] // remove last element
	fmt.Println(slice)
	slice = append(slice[:2], slice[3:]...) // remove 3 element
	fmt.Println("After 3rd remove : ", slice)
	slice = []string{"dog", "cat", "camel", "elephant", "lion", "tiger"}
	slicecopy := slices.Clone(slice)
	slices.Sort(slicecopy)
	fmt.Println("sorted : ", slicecopy)
	slices.Reverse(slicecopy)
	fmt.Println("Reversed : ", slicecopy)
	slicecopy = slices.Replace(slicecopy, 0, 1, "tigeress")
	fmt.Println(slicecopy)

	playstring()
	mapvars()
	structvars()
	httpJsonPlay()
}

func httpJsonPlay() {
	fmt.Println("############ Http & JSON playground #############")
	var word string
	fmt.Println("Enter word to look for meaning ?")
	fmt.Scan(&word)
	url := fmt.Sprintf("https://api.dictionaryapi.dev/api/v2/entries/en/%v", word)
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	var data []map[string]interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println(err)
	}
	meanings, _ := data[0]["meanings"].([]interface{})
	meaningsarr, _ := meanings[0].(map[string]interface{})
	defintionarr, _ := meaningsarr["definitions"].([]interface{})
	defintion, _ := defintionarr[0].(map[string]interface{})
	fmt.Println(data[0]["word"], " ====> ", defintion["definition"])

}

func playstring() {
	fmt.Println("############ String playground #############")
	var str = "hello world"
	arr := strings.Split(str, "")
	fmt.Printf("%v = %T \n", arr, arr)
	slices.Reverse(arr)
	fmt.Println(arr)
	arrnew := strings.Join(arr, "")
	fmt.Println(arrnew)
}

func structvars() {
	fmt.Println("############ Structs #############")
	type varstruct struct {
		name string
		age  int
	}
	a := varstruct{
		name: "faisal",
		age:  30,
	}
	fmt.Println(a.name)

}

func mapvars() {
	fmt.Println("############ Maps #############")
	var mapvar = make(map[int]string)

	mapvar[0] = "hello" // adding key values to map
	mapvar[1] = "world"
	mapvar[2] = "faisal"

	fmt.Println(mapvar)

	delete(mapvar, 0) // deleting any key
	fmt.Println("after delete : ", mapvar)

	mapvar2 := make(map[string]string)
	mapvar2["a"] = "aaa"
	mapvar2["b"] = "bbb"
	mapvar2["c"] = "ccc"
	fmt.Println(mapvar2)

	delete(mapvar2, "b")
	fmt.Println(mapvar2)

	var mapvar3 = map[string]string{} // defining using map only not making use of make
	mapvar3["a"] = "aaa"
	mapvar3["b"] = "bbb"
	mapvar3["c"] = "ccc"
	fmt.Println(mapvar3)

	delete(mapvar3, "b")
	fmt.Println(mapvar3)

}

func callbyreference(vartochange *int) {
	fmt.Println("############ callbyreference #############")
	*vartochange = 20
}

func callbyvalue(localvar2 string) {
	fmt.Println("############ callbyvalue #############")
	var localvar = "localtest"
	fmt.Println(globalvar)
	fmt.Println(localvar)
	fmt.Println(localvar2)
}
