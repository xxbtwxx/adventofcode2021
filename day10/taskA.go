package day10

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	Cmd.AddCommand(&cobra.Command{
		Use:   "a",
		Short: "Solution for day 10 problem A",
		Run: func(cmd *cobra.Command, args []string) {
			taskA()
		},
	})
}

func taskA() {
	value := 0

	for line := range inputReader() {
		inputStack := make([]rune, 0, len(line))
	inner:
		for _, el := range line {
			switch el {
			case '(', '[', '{', '<':
				inputStack = append(inputStack, el)
			case ')', ']', '}', '>':
				if inputStack[len(inputStack)-1] != bracketMatch[el] {
					value += valueMap[el]
					break inner
				}

				inputStack = inputStack[:len(inputStack)-1]
			default:
				fmt.Println("unexpected value", el)
				return
			}
		}
	}

	fmt.Println(value)
}
