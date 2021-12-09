package main

import (
	"adventofcode2021/day1"
	"adventofcode2021/day2"
	"adventofcode2021/day3"
	"adventofcode2021/day4"
	"adventofcode2021/day5"
	"adventofcode2021/day6"
	"adventofcode2021/day7"
	"adventofcode2021/day8"
	"adventofcode2021/day9"
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
	rootCmd.AddCommand(day5.Cmd)
	rootCmd.AddCommand(day6.Cmd)
	rootCmd.AddCommand(day7.Cmd)
	rootCmd.AddCommand(day8.Cmd)
	rootCmd.AddCommand(day9.Cmd)
}

func main() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
