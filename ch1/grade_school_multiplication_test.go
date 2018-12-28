package ch1

import "testing"

func TestMultiplication(t *testing.T) {
	a := "5678"
	b := "1234"

	want := "7006652"
	got, _ := GradeSchoolMultiply(a, b)

	if got != want {
		t.Errorf("want %s got %s", want, got)
	}
}
