package main

//斐波那契
func main() {

}

func fib(n int) int {
	if n < 2 {
		return n
	}

	a, b, c := 0, 1, 0
	for i := 1; i < n; i++ {
		c = a + b
		a, b = b, c
	}

	return c
}

func fib02(n int) int {
	a, b := 0, 1
	for i := 0; i < n; i++ {
		a, b = b, a+b
	}

	return a
}
