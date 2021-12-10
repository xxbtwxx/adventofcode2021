package day10

import "github.com/spf13/cobra"

var Cmd = &cobra.Command{
	Use:   "day10",
	Short: "Solutions for day10",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}
