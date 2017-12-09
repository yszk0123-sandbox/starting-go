package main

import "fmt"

func sum(s ...int) int {
	result := 0

	for _, v := range s {
		result += v
	}

	return result
}

func main() {
	fmt.Printf("%v\n", sum(10))
	fmt.Printf("%v\n", sum(10, 2))
	fmt.Printf("%v\n", sum(10, 2, 3))
}
