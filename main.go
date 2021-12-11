package main

import (
	"adventofcode2021/day01"
	"adventofcode2021/day02"
	"adventofcode2021/day03"
	"adventofcode2021/day04"
	"adventofcode2021/day05"
	"adventofcode2021/day06"
	"adventofcode2021/day07"
	"adventofcode2021/day08"
	"adventofcode2021/day09"
	"adventofcode2021/day10"
	"adventofcode2021/day11"
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
	rootCmd.AddCommand(day01.Cmd)
	rootCmd.AddCommand(day02.Cmd)
	rootCmd.AddCommand(day03.Cmd)
	rootCmd.AddCommand(day04.Cmd)
	rootCmd.AddCommand(day05.Cmd)
	rootCmd.AddCommand(day06.Cmd)
	rootCmd.AddCommand(day07.Cmd)
	rootCmd.AddCommand(day08.Cmd)
	rootCmd.AddCommand(day09.Cmd)
	rootCmd.AddCommand(day10.Cmd)
	rootCmd.AddCommand(day11.Cmd)
}

func main() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
