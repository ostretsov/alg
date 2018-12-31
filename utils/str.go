package utils

func StrSimpleReverse(input string) (output string) {
	for i := len(input) - 1; i >= 0; i-- {
		output = output + string(input[i])
	}

	return
}
