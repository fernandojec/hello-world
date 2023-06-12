package main

import "testing"

func TestGenerateWord(t *testing.T) {

	t.Run("Should Return Buzz", func(t *testing.T) {
		if got := GenerateWord(5); got == "Buzz" {
			t.Errorf("GenerateWord() = %v, want %v", got, "Buzz")
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
