package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	t, err := time.Parse(time.RFC3339, "2017-09-18T15:00:00+09:00")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(t)
}
