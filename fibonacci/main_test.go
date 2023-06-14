package main

import "testing"

func Test_fib(t *testing.T) {
	type args struct {
		i int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Should return 0",
			args: args{
				i: 0,
			},
			want: 0,
		}, {
			name: "Should return 1",
			args: args{
				i: 1,
			},
			want: 1,
		}, {
			name: "Should return 1",
			args: args{
				i: 2,
			},
			want: 1,
		}, {
			name: "Should return 2",
			args: args{
				i: 3,
			},
			want: 2,
		}, {
			name: "Should return 3",
			args: args{
				i: 4,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fib(tt.args.i); got != tt.want {
				t.Errorf("fib() = %v, want %v", got, tt.want)
			}
		})
	}
}
