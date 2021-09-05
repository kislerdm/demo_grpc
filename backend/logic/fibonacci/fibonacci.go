package fibonacci

import "math"

// Fibonacci calculates Fibonacci numbers.
func Fibonacci(n int64) int64 {
	if n < 0 {
		return Fibonacci(-n) * int64(math.Pow(-1, float64(-n+1)))
	}
	if n < 2 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}
