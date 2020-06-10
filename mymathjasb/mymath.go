// Package mymath provides math solutions
package mymath

// Sum suma un número ilimitado valores de tipo int
// Devuelve un valor de tipo int
func Sum(xi ...int) int {
	sum := 0
	for _, v := range xi {
		sum += v
	}
	return sum
}

// Average calcula el promedio de un número ilimitado de valores de tipo int
// Devuelve un valor de tipo float64
func Average(xi ...int) float64 {
	sum := 0
	for _, v := range xi {
		sum += v

	}
	return float64(sum / len(xi))
}
