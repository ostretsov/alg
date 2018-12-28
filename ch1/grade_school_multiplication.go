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

func ConvertStringNumberToBytes(number string) []byte {
	var result []byte
	for i := 0; i < len(number); i++ {
		b := byte(number[i]) - byte('0')
		result = append(result, b)
	}

	return result
}
