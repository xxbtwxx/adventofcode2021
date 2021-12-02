package day2

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	Cmd.AddCommand(&cobra.Command{
		Use:   "b",
		Short: "Solution for day 2 problem B",
		Run: func(cmd *cobra.Command, args []string) {
			taskB()
		},
	})
}

func taskB() {
	forward, depth, aim := 0, 0, 0

	commands := commandReader()

	for command := range commands {
		switch command.direction {
		case "forward":
			forward += command.value
			depth += command.value * aim
		case "up":
			aim -= command.value
		case "down":
			aim += command.value
		default:
			fmt.Println("invalid command direction")
			return
		}
	}

	fmt.Println(forward * depth)
}
