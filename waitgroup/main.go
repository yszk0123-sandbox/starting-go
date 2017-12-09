package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := new(sync.WaitGroup)
	wg.Add(2)

	go func() {
		for i := 0; i < 100; i++ {
			fmt.Print("1")
		}
		wg.Done()
	}()

	go func() {
		for i := 0; i < 100; i++ {
			fmt.Print("2")
		}
		wg.Done()
	}()

	wg.Wait()
}
