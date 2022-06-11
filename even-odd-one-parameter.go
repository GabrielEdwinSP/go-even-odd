package goevenodd

func EvenOdd(number int) string {
	var result string
	if number%2 == 0 {
		result = "Genap"
	} else {
		result = "Ganjil"
	}
	return result
}
