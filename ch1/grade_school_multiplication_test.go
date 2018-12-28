package ch1

import "testing"

func TestGradeSchoolMultiply(t *testing.T) {
	t.Run("invalid arguments", func(t *testing.T) {
		a := "5678.1"
		b := "1234"

		_, err := GradeSchoolMultiply(a, b)

		if err == nil {
			t.Errorf("expected error on invalid argument")
		}
	})

	t.Run("short numbers of same length", func(t *testing.T) {
		a := "5678"
		b := "1234"

		want := "7006652"
		got, _ := GradeSchoolMultiply(a, b)

		if got != want {
			t.Errorf("want %s got %s", want, got)
		}
	})

	t.Run("long numbers of same length", func(t *testing.T) {
		a := "5678293874908729034090283478162094598729348394875987123987234598701083548734909038857698911"
		b := "1234489750984985098037129384738479193757692173718457766635242173950696712902048171635493875"

		want := "7009795591655643027172461295023953881167542446640054000765770377106022505948006729526866614890982359431931298734376531996538221533914172897553259381603807010934556812446435534670125"
		got, err := GradeSchoolMultiply(a, b)

		if err != nil {
			t.Fatal(err)
		}

		if got != want {
			t.Errorf("want %s got %s", want, got)
		}
	})
}
