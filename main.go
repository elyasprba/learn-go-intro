package main

import (
	"fmt"

	Basic "github.com/elyasprba/learn-go-intro/basic"
	// Intro "github.com/elyasprba/learn-go-intro/intro"
	Server "github.com/elyasprba/learn-go-intro/server"
)

func main() {
	fmt.Println("Hello World! my name is El. I am learning Go!")

	// INTRO
	// Intro.Conditionals()
	// Intro.Variabels()
	// Intro.Conditionals()
	// Intro.Looping()
	// Intro.Arrs()
	// Intro.Maps()
	// Intro.Login("El", "1234")
	// var a, b = I.CheckAge(2000)
	// fmt.Println(a, b)

	// BASIC
	Basic.Main()

	// SERVER
	Server.Routes()

}
