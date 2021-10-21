package projecteuler

import (
	"testing"
)

func TestAnswer(t *testing.T) {
	cases := []struct {
		limit int
		sum   int
	}{
		{10, 23},
		{1000, 233168},
	}

	for _, c := range cases {
		got := answer(c.limit)
		want := c.sum

		if got != c.sum {
			t.Errorf("answer() = %d; want = %d", got, want)
		}
	}
}
