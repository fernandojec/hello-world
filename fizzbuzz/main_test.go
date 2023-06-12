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

	// auto generate
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "Should Return Fizz",
			args: args{
				n: 3,
			},
			want: "Fizz",
		},
		{
			name: "Should Return Buzz",
			args: args{
				n: 5,
			},
			want: "Buzz",
		}, {
			name: "Should Return FizzBuzz",
			args: args{
				n: 15,
			},
			want: "FizzBuzz",
		}, {
			name: "Should Return 4",
			args: args{
				n: 4,
			},
			want: "4",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GenerateWord(tt.args.n); got != tt.want {
				t.Errorf("GenerateWord() = %v, want %v", got, tt.want)
			}
		})
	}
}
