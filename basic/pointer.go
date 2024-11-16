package basic

import (
	"fmt"
)

func Pointer() {
	var a int = 10
	var b *int = &a

	fmt.Println(a)
	fmt.Println(b)
	// fmt.Println(*b)
}
