package main

import (
	"adventofcode2021/day1"
	"adventofcode2021/day2"
	"adventofcode2021/day3"
	"adventofcode2021/day4"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = cobra.Command{
	Use:   "adventofcode",
	Short: "Advent of code 2021 solutions",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	rootCmd.AddCommand(day1.Cmd)
	rootCmd.AddCommand(day2.Cmd)
	rootCmd.AddCommand(day3.Cmd)
	rootCmd.AddCommand(day4.Cmd)
}

func main() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
