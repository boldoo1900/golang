package fizzbuzz

import (
	"fmt"
	"testing"
)

func FizzBuzz(i int) string {

	if i%15 == 0 {
		// fmt.Println("FizzBuzz")
		return "FizzBuzz"
	} else if i%3 == 0 {
		// fmt.Println("Fizz")
		return "Fizz"
	} else if i%5 == 0 {
		// fmt.Println("Buzz")
		return "Buzz"
	} else {
		// fmt.Println(i)
		return string(i)
	}
}
func TestFizzBuzzIndex(t *testing.T) {
	tests := []struct {
		name string
		x    int
		want string
	}{
		{name: "case1", x: 3, want: "Fizz"},
		{name: "case2", x: 5, want: "Buzz"},
		{name: "case3", x: 15, want: "FizzBuzz"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FizzBuzz(tt.x); got != tt.want {
				fmt.Println(tt.x, got)
				t.Errorf("FizzBuzz(%v) = %v, want %v", tt.x, got, tt.want)
			}
		})
	}
}

func BenchmarkFizzBuzz(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FizzBuzz(5)
	}
}
