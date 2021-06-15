package fibo

// Seq returns a Fibonacci's sequense
func Seq(n int) int {
	x0, x1 := 0, 1
	for i := 1; i <= n; i++ {
		x0, x1 = x1, x0+x1
	}
	return x1
}
