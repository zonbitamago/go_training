package calc

import (
	"go_training/04_test/testutils"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	println("before all...")

	code := m.Run()

	println("after all...")

	os.Exit(code)
}

func TestSum(t *testing.T) {
	t.Run("sum_plus", func(t *testing.T) {
		want := 3
		if got := Sum(1, 2); got != want {
			testutils.ErrorfHandler(t, want, got)
		}
	})

	t.Run("sum_minus", func(t *testing.T) {
		want := -3
		if got := Sum(-1, -2); got != want {
			testutils.ErrorfHandler(t, want, got)
		}

	})
}

func TestSubtract(t *testing.T) {
	t.Run("subtract_plus", func(t *testing.T) {
		want := 2
		if got := Subtruct(5, 3); got != want {
			testutils.ErrorfHandler(t, want, got)
		}
	})

	t.Run("subtract_minus", func(t *testing.T) {
		want := -2
		if got := Subtruct(-5, -3); got != want {
			testutils.ErrorfHandler(t, want, got)
		}
	})

}
func TestMultiple(t *testing.T) {
	test := []struct {
		name string
		a    int
		b    int
		want int
	}{
		{"plus_pluse", 2, 3, 6},
		{"plus_minus", 3, -4, -12},
		{"minus_minus", -4, -5, 20},
	}

	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			if got := Multiple(tt.a, tt.b); got != tt.want {
				testutils.ErrorfHandler(t, tt.want, got)
			}
		})
	}
}
