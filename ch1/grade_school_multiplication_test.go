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

	want := []byte{0, 4, 3, 2, 1}
	got := convertStringNumberToBytes(num)

	if !reflect.DeepEqual(want, got) {
		t.Errorf("want %v got %v", want, got)
	}
}

func TestMultiplyByteByByte(t *testing.T) {
	a := []byte{1, 2, 3, 4}
	b := []byte{7, 8, 9}

	want := cascade{
		{7, 4, 2, 0, 3},
		{0, 8, 6, 5, 4, 3},
		{0, 0, 9, 8, 8, 8, 3},
	}
	got := multiplyByteByByte(a, b)

	if !reflect.DeepEqual(want, got) {
		t.Errorf("want %v, got %v", want, got)
	}
}

func TestMultiplyBytes(t *testing.T) {
	cases := []struct {
		Name                    string
		A, B                    byte
		WantReminder, WantCarry byte
	}{
		{"regular", byte(9), byte(9), byte(1), byte(8)},
		{"zeros", byte(0), byte(0), byte(0), byte(0)},
		{"zero remainder", byte(5), byte(4), byte(0), byte(2)},
		{"zero carry", byte(2), byte(1), byte(2), byte(0)},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			a := test.A
			b := test.B

			wantRemainder := test.WantReminder
			wantCarry := test.WantCarry
			gotRemainder, gotCarry := multiplyBytes(a, b)

			if gotRemainder != wantRemainder {
				t.Errorf("got %v want %v, case %v", gotRemainder, wantRemainder, test)
			}

			if gotCarry != wantCarry {
				t.Errorf("got %v want %v, case %v", gotCarry, wantCarry, test)
			}
		})
	}
}

func TestInitRow(t *testing.T) {
	want := cascade{
		{},
		{0},
	}
	got := cascade{}
	got.initRow(1)

	if !reflect.DeepEqual(want, got) {
		t.Errorf("got %v, want %v", got, want)
	}
}
