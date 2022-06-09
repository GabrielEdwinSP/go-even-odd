package goevenodd

import "fmt"

func EvenOdd(number int) {
	if number%2 == 0 {
		fmt.Println("Genap")
	} else {
		fmt.Println("Ganjil")
	}
}
