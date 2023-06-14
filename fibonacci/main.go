package main

import "fmt"

func main() {
	fmt.Println(fib(0))
	fmt.Println(fib(1))
	fmt.Println(fib(2))
	fmt.Println(fib(3))
	fmt.Println(fib(4))
}

func fib(i int) int {
	if i < 2 {
		return i
	}

	return fib(i-1) + fib(i-2)
}
