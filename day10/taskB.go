package day10

import (
	"fmt"
	"sort"

	"github.com/spf13/cobra"
)

func init() {
	Cmd.AddCommand(&cobra.Command{
		Use:   "b",
		Short: "Solution for day10 problem B",
		Run: func(cmd *cobra.Command, args []string) {
			taskB()
		},
	})
}

func taskB() {
	lineValues := []int{}
	incompleteLines := []string{}

	for line := range inputReader() {
		inputStack := make([]rune, 0, len(line))
		incompleteLine := true
		for _, el := range line {
			switch el {
			case '(', '[', '{', '<':
				inputStack = append(inputStack, el)
			case ')', ']', '}', '>':
				if inputStack[len(inputStack)-1] != bracketMatch[el] {
					incompleteLine = false
				}

				inputStack = inputStack[:len(inputStack)-1]
			default:
				fmt.Println("unexpected value", el)
				return
			}
		}

		if incompleteLine {
			incompleteLines = append(incompleteLines, line)
		}
	}

	for _, line := range incompleteLines {
		lineValue := 0
		lineStack := make([]rune, 0, len(line))
		for _, el := range line {
			if _, ok := openingBrackets[el]; ok {
				lineStack = append(lineStack, el)
			} else {
				lineStack = lineStack[:len(lineStack)-1]
			}
		}

		for len(lineStack) != 0 {
			lineValue *= 5
			lineValue += valueMap[lineStack[len(lineStack)-1]]
			lineStack = lineStack[:len(lineStack)-1]
		}

		lineValues = append(lineValues, lineValue)
	}

	sort.Ints(lineValues)

	fmt.Println(lineValues[len(lineValues)/2])

}
