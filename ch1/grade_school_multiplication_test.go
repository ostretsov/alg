package ch1

import (
	"reflect"
	"testing"
)

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

func TestConvertStringNumberToBytes(t *testing.T) {
	num := "12340"

	want := []byte{1, 2, 3, 4, 0}
	got := convertStringNumberToBytes(num)

	if !reflect.DeepEqual(want, got) {
		t.Errorf("want %v got %v", want, got)
	}
}

//func TestMultiplyByteByByte(t *testing.T) {
//	a := []byte{1, 2, 3, 4}
//	b := []byte{7, 8, 9}
//
//	want := cascade{
//		{1, 1, 1, 0, 6},
//		{9, 8, 7, 2, 0},
//		{8, 6, 3, 8, 0, 0},
//	}
//	got := multiplyByteByByte(a, b)
//
//	if (!reflect.DeepEqual(want, got)) {
//		t.Errorf("want %v, got %v", want, got)
//	}
//}

func TestMultiplyBytes(t *testing.T) {
	cases := []struct {
		A, B                    byte
		WantReminder, WantCarry byte
	}{
		{byte(9), byte(9), byte(1), byte(8)},
		{byte(0), byte(0), byte(0), byte(0)},
		{byte(5), byte(4), byte(0), byte(2)},
	}

	for _, c := range cases {
		a := c.A
		b := c.B

		wantRemainder := c.WantReminder
		wantCarry := c.WantCarry
		gotRemainder, gotCarry := multiplyBytes(a, b)

		if gotRemainder != wantRemainder {
			t.Errorf("got %v want %v, case %v", gotRemainder, wantRemainder, c)
		}

		if gotCarry != wantCarry {
			t.Errorf("got %v want %v, case %v", gotCarry, wantCarry, c)
		}
	}

}
