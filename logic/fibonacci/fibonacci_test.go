package fibonacci_test

import (
	"fmt"
	"testing"

	"github.com/kislerdm/demo-grcp/logic/fibonacci"
)

var Tests = []struct {
	in   int64
	want int64
}{
	{
		in:   0,
		want: 0,
	},
	{
		in:   1,
		want: 1,
	},
	{
		in:   2,
		want: 1,
	},
	{
		in:   3,
		want: 2,
	},
	{
		in:   5,
		want: 5,
	},
	{
		in:   6,
		want: 8,
	},
	{
		in:   20,
		want: 6765,
	},
	{
		in:   -1,
		want: 1,
	},
	{
		in:   -20,
		want: -6765,
	},
	{
		in:   -5,
		want: 5,
	},
}

func TestFibonacci(t *testing.T) {
	for _, test := range Tests {
		test := test
		t.Run(fmt.Sprintf("%d", test.in), func(t *testing.T) {
			t.Parallel()
			got := fibonacci.Fibonacci(test.in)
			if test.want != got {
				t.Fatalf("\nin: %v\nwant: %v\ngot: %v", test.in, test.want, got)
			}
		})
	}
}
