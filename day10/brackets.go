package day10

var valueMap = map[rune]int{
	')': 3,
	']': 57,
	'}': 1197,
	'>': 25137,

	'(': 1,
	'[': 2,
	'{': 3,
	'<': 4,
}

var bracketMatch = map[rune]rune{
	')': '(',
	']': '[',
	'}': '{',
	'>': '<',

	'(': ')',
	'[': ']',
	'{': '}',
	'<': '>',
}

var openingBrackets = map[rune]struct{}{
	'(': {},
	'[': {},
	'{': {},
	'<': {},
}
