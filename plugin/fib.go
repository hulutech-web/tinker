package plugin

func Fib(n int) int {
	return fib(n, 0, 1)
}

func fib(n, a, b int) int {
	if n == 0 {
		return a
	} else if n == 1 {
		return b
	}
	return fib(n-1, b, a+b)
}
