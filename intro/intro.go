package intro

import "fmt"

func Variabels() {
	// manifest type
	var firstName string = "El"
	const lastName string = "dut"

	fmt.Println(firstName + lastName)

	// interface type
	hobby := "Coding"
	fmt.Println(hobby)
}

func Conditionals() {
	var name string = "El"
	var password string = "1234"

	if name == "El" && password == "1234" {
		fmt.Println("You are El")
	} else {
		fmt.Println("You are not El")
	}
}

func Looping() {
	// case 1
	for i := 0; i < 10; i++ {
		fmt.Println("angka ke", i)
	}

	// case 2

	var number int = 0

	for {
		if number == 10 {
			break
		}

		fmt.Println("number ke", number)

		number++
	}
}

func Arrs() {

	// Array
	var fruits = [3]string{"apple", "banana", "grape"}

	// for i := 0; i < len(fruits); i++ {
	// 	fmt.Println(fruits[i])
	// }

	for _, fruit := range fruits {
		fmt.Println(fruit)
	}

	// Slice
	var numbers = []int{1, 2, 3, 4, 5}

	for _, number := range numbers {
		fmt.Println(number)
	}
}

func Maps() {
	var person = map[string]string{
		"name":    "El",
		"address": "Jakarta",
		"age":     "20",
	}

	for key, value := range person {
		fmt.Println(key, value)
	}
}

func Login(user, password string) {
	if user == "El" && password == "1234" {
		fmt.Println("Success Login")
	} else {
		fmt.Println("Username or Password is wrong!")
	}
}

func CheckAge(year int) (bool, int) {

	if age := 2024 - year; age < 18 {
		return false, age
	} else {
		return true, age
	}

}
