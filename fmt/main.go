package main

import (
	"fmt"
)

func main() {
	fmt.Printf("%d, %d\n", 10)
	fmt.Printf("%d, %d\n", 10, 20)
	fmt.Printf("%d, %d\n", 10, 20, 30)
	fmt.Printf("%v\n", []int{1, 2, 3})
	fmt.Printf("%#v\n", []int{1, 2, 3})
	fmt.Printf("%T\n", []int{1, 2, 3})
}
