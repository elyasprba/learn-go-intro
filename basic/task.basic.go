package basic

import "fmt"

type Deret struct {
	Prima     int
	Ganjil    int
	Genap     int
	Fibonacci int
}

func TaskMethod() {
	var deret = Deret{
		Prima:     40,
		Ganjil:    40,
		Genap:     40,
		Fibonacci: 40,
	}

	fmt.Println("Bilangan prima hingga", deret.Prima, ":")
	deret.PrintPrima()
	fmt.Println()

	fmt.Println("Bilangan ganjil hingga", deret.Ganjil, ":")
	deret.PrintGanjil()
	fmt.Println()

	fmt.Println("Bilangan genap hingga", deret.Genap, ":")
	deret.PrintGenap()
	fmt.Println()

	fmt.Println("Bilangan Fibonacci hingga", deret.Fibonacci, ":")
	deret.PrintFibonacci()
	fmt.Println()
}

// Fungsi untuk mencetak bilangan prima
func (d *Deret) PrintPrima() {
	for i := 2; i <= d.Prima; i++ {
		isPrime := true
		for j := 2; j*j <= i; j++ {
			if i%j == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			fmt.Print(i, " ")
		}
	}
}

// Fungsi untuk mencetak bilangan ganjil
func (d *Deret) PrintGanjil() {
	for i := 1; i <= d.Ganjil; i += 2 {
		fmt.Print(i, " ")
	}
}

// Fungsi untuk mencetak bilangan genap
func (d *Deret) PrintGenap() {
	for i := 2; i <= d.Genap; i += 2 {
		fmt.Print(i, " ")
	}
}

// Fungsi untuk mencetak bilangan Fibonacci
func (d *Deret) PrintFibonacci() {
	a, b := 0, 1
	for a <= d.Fibonacci {
		fmt.Print(a, " ")
		a, b = b, a+b
	}
}
