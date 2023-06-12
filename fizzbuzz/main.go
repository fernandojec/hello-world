package main

import "fmt"

func main() {
	for n := 1; n <= 15; n++ {
		GenerateWord(n)
	}
}

func GenerateWord(n int) string {
	if n%3 == 0 && n%5 == 0 {
		return "FizzBuzz"
	} else if n%3 == 0 {
		return "Fizz"
	} else if n%5 == 0 {
		return "Buzz"
	} else {
		return fmt.Sprintf("%d", n)
	}
}
