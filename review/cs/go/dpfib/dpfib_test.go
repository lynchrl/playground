package dpfib

import (
	"testing"
)

func TestFibSlow(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{
			name: "Fibonacci of 1",
			n:    1,
			want: 1,
		},
		{
			name: "Fibonacci of 2",
			n:    2,
			want: 1,
		},
		{
			name: "Fibonacci of 3",
			n:    3,
			want: 2,
		},
		{
			name: "Fibonacci of 4",
			n:    4,
			want: 3,
		},
		{
			name: "Fibonacci of 5",
			n:    5,
			want: 5,
		},
		{
			name: "Fibonacci of 10",
			n:    10,
			want: 55,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fibSlow(tt.n); got != tt.want {
				t.Errorf("fibSlow() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFibMemo(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{
			name: "Fibonacci of 1",
			n:    1,
			want: 1,
		},
		{
			name: "Fibonacci of 2",
			n:    2,
			want: 1,
		},
		{
			name: "Fibonacci of 3",
			n:    3,
			want: 2,
		},
		{
			name: "Fibonacci of 4",
			n:    4,
			want: 3,
		},
		{
			name: "Fibonacci of 5",
			n:    5,
			want: 5,
		},
		{
			name: "Fibonacci of 10",
			n:    10,
			want: 55,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fibMemo(tt.n); got != tt.want {
				t.Errorf("fibMemo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkFibSlow(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fibSlow(20)
	}
}

func BenchmarkFibMemo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fibMemo(20)
	}
}
