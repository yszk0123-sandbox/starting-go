package main

import (
	"fmt"
)

func main() {
	{
		fmt.Printf("Copy s2 to s1 (s1 > s2)\n")
		s1 := []int{1, 2, 3, 4, 5}
		s2 := []int{10, 11}
		n := copy(s1, s2)
		fmt.Printf("%v, %v, %v\n", s1, s2, n)
	}
	{
		fmt.Printf("Copy s2 to s1 (s1 < s2)\n")
		s1 := []int{1, 2, 3}
		s2 := []int{10, 11, 12, 13, 14}
		n := copy(s1, s2)
		fmt.Printf("%v, %v, %v\n", s1, s2, n)
	}
}
