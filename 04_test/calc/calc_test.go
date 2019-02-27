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
