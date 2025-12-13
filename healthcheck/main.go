package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func main() {
	for {
		time.Sleep(5 * time.Second)
		checkStatus()
	}
}

func checkStatus() {
	urls := []string{
		"www.google.com",
		"www.yahoo.com",
		"www.msn.com",
	}
	ch := make(chan string, 3)
	for _, v := range urls {
		go func() {
			resp, err := http.Get(v)
			if err != nil {
				// handle error
			}
			defer resp.Body.Close()
			body, err := io.ReadAll(resp.Body)
			// ch <- body.statusCode
			fmt.Println(body)
			ch <- "a"
		}()
	}

	for v := range ch {
		fmt.Println("status : %v", v)
	}

}
