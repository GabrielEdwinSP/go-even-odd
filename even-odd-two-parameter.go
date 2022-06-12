package goevenodd

func EvenOddTwo(number ...int) []string {
	result := []string{}

	for _, i := range number {
		if i%2 == 0 {
			result = append(result, "Genap")
		} else {
			result = append(result, "Ganjil")
		}
	}
	return result
}
