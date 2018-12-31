package ch1

import (
	"alg/utils"
	"fmt"
	"strconv"
)

type cascade [][]byte

func (c *cascade) initRow(n int) {

}

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
	/**
	if len(b) > len(a) {
		a, b = b, a
	}
	bytesA := convertStringToBytes(a)
	bytesB := convertStringToBytes(b)
	c := multiplyByteByByte(a, b)
	return sumCascadeWithCarries(c)
	*/
}

func convertStringNumberToBytes(number string) []byte {
	number = utils.StrSimpleReverse(number)

	var result []byte
	for i := 0; i < len(number); i++ {
		result = append(result, number[i]-'0')
	}

	return result
}

func multiplyByteByByte(a, b []byte) (c cascade) {
	var carry byte = 0
	for kB, vB := range b {
		c.initRow(kB)

		for _, vA := range a {
			r, cr := multiplyBytes(vA, vB)
			c[kB] = append(c[kB], r+carry)
			carry = cr
		}
	}

	return
}

func multiplyBytes(a, b byte) (remainder, carry byte) {
	product := a * b
	remainder = product % 10
	carry = (product - remainder) / 10

	return
}
