package day08

import "github.com/spf13/cobra"

var Cmd = &cobra.Command{
	Use:   "day8",
	Short: "Solutions for day 8",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}
