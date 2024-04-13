package calculations

import (
	"testing"
)

func TestDoMaths(t *testing.T) {
	t.Run("Test adding two integers", func(t *testing.T) {
		got := DoMaths(1, 2, "+")
		want := 3
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
	t.Run("test multiplying two integers", func(t *testing.T) {
		got := DoMaths(4, 4, "x")
		want := 16
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
	t.Run("test dividing two integers", func(t *testing.T) {
		got := DoMaths(4, 4, "/")
		want := 1
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
	t.Run("test subtracting two integers", func(t *testing.T) {
		got := DoMaths(42, 4, "-")
		want := 38
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

}
