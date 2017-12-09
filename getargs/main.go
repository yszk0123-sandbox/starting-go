package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf("Number of args: %d\n", len(os.Args))
	for i, arg := range os.Args {
		fmt.Printf("Arg[%d] = %v\n", i, arg)
	}
}
