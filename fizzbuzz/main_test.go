package main

import "testing"

func TestGenerateWord(t *testing.T) {

	t.Run("Should Return Buzz", func(t *testing.T) {
		if got := GenerateWord(5); got != "Buzz" {
			t.Errorf("GenerateWord() = %v, want %v", got, "Buzz")
		}
	})
	t.Run("Should Return Fizz", func(t *testing.T) {
		if got := GenerateWord(3); got != "Fizz" {
			t.Errorf("GenerateWord() = %v, want %v", got, "Fizz")
		}
	})
	t.Run("Should Return FizzBuzz", func(t *testing.T) {
		if got := GenerateWord(15); got != "FizzBuzz" {
			t.Errorf("GenerateWord() = %v, want %v", got, "FizzBuzz")
		}
	})
	t.Run("Should Return 4", func(t *testing.T) {
		if got := GenerateWord(4); got != "4" {
			t.Errorf("GenerateWord() = %v, want %v", got, "4")
		}
	})
	// type args struct {
	// 	n int
	// }
	// tests := []struct {
	// 	name string
	// 	args args
	// 	want string
	// }{
	// 	// TODO: Add test cases.
	// }
	// for _, tt := range tests {
	// 	t.Run(tt.name, func(t *testing.T) {
	// 		if got := GenerateWord(tt.args.n); got != tt.want {
	// 			t.Errorf("GenerateWord() = %v, want %v", got, tt.want)
	// 		}
	// 	})
	// }
}
