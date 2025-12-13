package main

import (
	"fmt"
	"sync"
	"time"
)

func DoWork(i int) int {

	time.Sleep(1 * time.Second)
	// return rand.Intn(100)
	return i
}

func main() {
	ch := make(chan int, 1000)
	go func() {
		wg := sync.WaitGroup{}
		for i := 1; i <= 1000; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				n := DoWork(i)
				ch <- n
				// close(ch)
			}()
		}
		wg.Wait()
		close(ch)
	}()

	// n1 := <-ch
	// n2 := <-ch

	// fmt.Println(n1, n2)
	for v := range ch {
		fmt.Println(v)
	}

}
