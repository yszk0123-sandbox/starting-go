package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id   int    "user id"
	Name string "user name"
}

func main() {
	user := User{Id: 1, Name: "foo"}
	fmt.Printf("%v\n", user)

	info := reflect.TypeOf(user)
	for i := 0; i < info.NumField(); i++ {
		field := info.Field(i)
		fmt.Printf("%d = %v (%v)\n", i, field.Name, field.Tag)
	}
}
