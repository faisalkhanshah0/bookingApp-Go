package main

import (
	"fmt"
	"strings"
)

func main() {

	// Basic Checks & Info
	strings.Contains("hello", "he")      // true
	strings.ContainsAny("team", "aeiou") // true
	strings.ContainsRune("golang", 'g')  // true

	strings.HasPrefix("hello", "he") // true
	strings.HasSuffix("hello", "lo") // true

	strings.Count("cheese", "e")         // 3
	strings.Index("chicken", "ken")      // 4
	strings.IndexAny("failure", "xyzai") // 1
	strings.IndexByte("golang", 'o')     // 1
	strings.IndexRune("golang", 'g')     // 0
	strings.LastIndex("go gopher", "go") // 3

	// Splitting & Joining
	strings.Split("a,b,c", ",")          // ["a" "b" "c"]
	strings.SplitN("a,b,c", ",", 2)      // ["a" "b,c"]
	strings.SplitAfter("a,b,c", ",")     // ["a," "b," "c"]
	strings.SplitAfterN("a,b,c", ",", 2) // ["a," "b,c"]

	strings.Fields(" foo  bar baz  ")     // ["foo" "bar" "baz"] (splits on whitespace)
	strings.Join([]string{"a", "b"}, "-") // "a-b"

	// Replacing
	strings.Replace("oink oink", "k", "ky", 1)     // "oinky oink"
	strings.ReplaceAll("oink oink", "oink", "moo") // "moo moo"

	strings.Repeat("na", 3) // "nanana"

	// Case Conversion
	strings.ToLower("GoLang")    // "golang"
	strings.ToUpper("GoLang")    // "GOLANG"
	strings.Title("hello world") // "Hello World" (deprecated, use cases.Title in unicode package)

	strings.EqualFold("Go", "go") // true (case-insensitive compare)

	// Trimming
	strings.TrimSpace("   hello   ")  // "hello"
	strings.Trim("!!hello!!", "!")    // "hello"
	strings.TrimLeft("xxhello", "x")  // "hello"
	strings.TrimRight("helloxx", "x") // "hello"

	strings.TrimPrefix("foo_bar", "foo_") // "bar"
	strings.TrimSuffix("foo_bar", "_bar") // "foo"

	// Modification
	strings.Map(func(r rune) rune {
		if r == 'a' {
			return 'A'
		}
		return r
	}, "banana") // "bAnAnA"

	strings.Clone("hello") // makes a new copy of string (new memory)

	// Cutting
	strings.Cut("a=b", "=") // "a", "b", true
	// If separator not found → returns whole string, "", false

	strings.CutPrefix("foo_bar", "foo_") // "bar", true
	strings.CutSuffix("foo_bar", "_bar") // "foo", true

	// Counting & Measuring
	strings.Compare("a", "b") // -1 (a < b)
	strings.Compare("a", "a") // 0
	strings.Compare("b", "a") // 1

	// strings.Builder{} // efficient string building
	strings.NewReader("hello") // *strings.Reader (implements io.Reader)

	// Reader & Builder Helpers
	// Using Builder (better than += in loops)
	var b strings.Builder
	b.WriteString("hello")
	b.WriteByte(' ')
	b.WriteString("world")
	fmt.Println(b.String()) // "hello world"

	// Using Reader
	r := strings.NewReader("GoLang")
	buf := make([]byte, 2)
	n, _ := r.Read(buf) // reads "Go"
	fmt.Println(n)

}
