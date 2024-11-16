package basic

import (
	"fmt"
)

type User struct {
	Name string
	Age  int
}

func Struct() {
	var user User = User{
		Name: "El",
		Age:  20,
	}

	// var user User
	// user.Name = "El"
	// user.Age = 20

	fmt.Println(`My name is`, user.Name, "and my age is", user.Age)
}
