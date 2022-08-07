package main

import "fmt"

func SumInts(m map[string]int64) int64 {
	var s int64
	for _, v := range m {
		s += v
	}
	return s
}

func SumFloats(m map[string]float64) float64 {
	var s float64
	for _, v := range m {
		s += v
	}
	return s
}

func SumIntsOrFloats[K comparable, V int64 | float64](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

func main() {
	ints := map[string]int64{
		"first":  34,
		"second": 12,
	}

	floats := map[string]float64{
		"first":  35.12,
		"second": 69.11,
	}

	int_ints := map[int64]int64{
		1: 2,
		3: 4,
	}

	fmt.Printf("Generic Sums: %v and %v and %v\n",
		SumIntsOrFloats(ints),
		SumIntsOrFloats(floats),
		SumIntsOrFloats(int_ints),
	)
}
