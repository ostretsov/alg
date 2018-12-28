package ch1

import (
	"fmt"
	"strconv"
)

func GradeSchoolMultiply(a, b string) (result string, err error) {
	intA, err := strconv.Atoi(a)
	if err != nil {
		return "", fmt.Errorf("can't convert %s to integer", a)
	}

	intB, err := strconv.Atoi(b)
	if err != nil {
		return "", fmt.Errorf("can't convert %s to integer", b)
	}

	return strconv.Itoa(intA * intB), nil
}
