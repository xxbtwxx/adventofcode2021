package day03

import "github.com/spf13/cobra"

var Cmd = &cobra.Command{
	Use:   "day3",
	Short: "Solutions for day 3",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}
