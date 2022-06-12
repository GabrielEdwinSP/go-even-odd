package goevenodd

import "testing"

func BenchmarkOneParameterTable(b *testing.B) {
	benchmarks := []struct {
		name    string
		request int
	}{
		{
			name:    "1",
			request: 1,
		},
		{
			name:    "2",
			request: 2,
		},
		{
			name:    "3",
			request: 3,
		},
	}
	for _, benbenchmark := range benchmarks {
		b.Run(benbenchmark.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				EvenOdd(benbenchmark.request)
			}
		})
	}
}

func BenchmarkTwoParameterTable(b *testing.B) {
	benchmarks := []struct {
		name    string
		request []int
	}{
		{
			name:    "1",
			request: []int{1, 2, 3},
		},
		{
			name:    "2",
			request: []int{3, 2, 5},
		},
		{
			name:    "3",
			request: []int{29, 231, 23, 22, 200},
		},
	}
	for _, benbenchmark := range benchmarks {
		b.Run(benbenchmark.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				EvenOddTwo(benbenchmark.request...)
			}
		})
	}
}
