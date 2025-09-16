package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {

	data, err := os.ReadFile("dummy.json")
	if err != nil {
		panic(err)
	}
	// var jsonobj any
	var jsonobj []map[string]any
	var educobj []any
	err = json.Unmarshal(data, &jsonobj)
	if err != nil {
		panic(err)
	}
	educobj = jsonobj[0]["education"].([]any)
	pretty, err := json.MarshalIndent(jsonobj, "", "  ")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(pretty))
	fmt.Println(jsonobj[0]["education"])
	fmt.Println(educobj[0])

	type structvar struct {
		name      string
		age       int
		education []any
	}

	var obj = []structvar{
		{"ram", 20, []any{"graduage"}},
		{"ramu", 21, []any{"graduage"}},
		{"shyam", 23, []any{"graduage"}},
	}
	fmt.Println(obj[0].name, obj[0].education[0])
	// fmt.Println(ram.education[0], shyam.education[2])

}
