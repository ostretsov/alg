package utils

import "testing"

func TestStrSimpleReverse(t *testing.T) {
	t.Run("empty string", func(t *testing.T) {
		want := ""
		got := StrSimpleReverse("")

		if got != want {
			t.Errorf("got %s want empty string", got)
		}
	})

	t.Run("regular string", func(t *testing.T) {
		want := "dcba"
		got := StrSimpleReverse("abcd")

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})
}
