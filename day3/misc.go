package day3

var runeToVal = map[rune]int{
	'0': 0,
	'1': 1,
}

func binaryIntSliceToInt(slice []int) int {
	value := 0
	for _, vals := range slice {
		value = value << 1
		value += vals
	}

	return value
}

func binaryStringToInt(s string) int {
	value := 0
	for _, vals := range s {
		value = value << 1
		value += runeToVal[vals]
	}

	return value
}
