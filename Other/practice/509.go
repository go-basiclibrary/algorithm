package main

// 509 斐波那契
func main() {

}

func fib(n int) int {
	// special situation
	if n <= 2 {
		return n
	}

	a, b, c := 0, 1, 0
	for i := 1; i < n; i++ {
		c = a + b
		a, b = b, c
	}

	return c
}
