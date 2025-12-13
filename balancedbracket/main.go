package main

import (
	"fmt"
	"strings"
	// "slices"
)

func main (){


// {{[}]}
// {[]{}}
// {[]{[}]}
// 12123210 

inputString :=  "{[]{}}()"

// setOfBraces := map[string]string{
// 	'(': ')',
// 	'{': '}',
// 	'[': ']',
// }

// arropen :=  []string {"(","{","["}
// arrclose :=  []string {"}","}","]"}

for _, k := range inputString {
	fmt.Println(k)
	ch := fmt.Sprintf("%c",k)
	fmt.Println(ch)
	if(strings.Contains(inputString, ch)){
		fmt.Println("found")
	} else{
		fmt.Println("not found")
	}
	
}





}