package factorial_test

import (
	"fmt"
	"testing"

	"backend/logic/factorial"
)

var Tests = []struct {
	in   int64
	want int64
}{
	{
		in:   0,
		want: 1,
	},
	{
		in:   1,
		want: 1,
	},
	{
		in:   2,
		want: 2,
	},
	{
		in:   3,
		want: 6,
	},
	{
		in:   5,
		want: 120,
	},
	{
		in:   6,
		want: 720,
	},
	{
		in:   -1,
		want: -1,
	},
	{
		in:   -20,
		want: -1,
	},
}

func TestFactorial(t *testing.T) {
	for _, test := range Tests {
		test := test
		t.Run(fmt.Sprintf("%d", test.in), func(t *testing.T) {
			t.Parallel()
			got := factorial.Factorial(test.in)
			if test.want != got {
				t.Fatalf("\nin: %v\nwant: %v\ngot: %v", test.in, test.want, got)
			}
		})
	}
}
