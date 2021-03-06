package day02

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	Cmd.AddCommand(&cobra.Command{
		Use:   "a",
		Short: "Solution for day 2 problem A",
		Run: func(cmd *cobra.Command, args []string) {
			taskA()
		},
	})
}

func taskA() {
	forward, depth := 0, 0

	for command := range commandReader() {
		switch command.direction {
		case "forward":
			forward += command.value
		case "up":
			depth -= command.value
		case "down":
			depth += command.value
		default:
			fmt.Println("invalid command direction")
			return
		}
	}

	fmt.Println(forward * depth)
}
