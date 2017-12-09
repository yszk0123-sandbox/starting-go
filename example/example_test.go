package examples

import (
	"fmt"
)

func Add(x, y int) int {
	return x + y
}

func ExampleAdd() {
	fmt.Println(Add(1, 2))
	// Output: 3
}
