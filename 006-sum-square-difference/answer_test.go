package projecteuler

import (
	"testing"
)

func TestAnswer(t *testing.T) {
	cases := []struct {
		limit      int
		difference int
	}{
		{10, 2640},
		{100, 25164150},
	}

	for _, c := range cases {
		got := answer(c.limit)
		want := c.difference

		if got != want {
			t.Errorf("answer(%d) = %d; want = %d", c.limit, got, want)
		}
	}
}

func BenchmarkAnswer(b *testing.B) {
	answer(10)
}
