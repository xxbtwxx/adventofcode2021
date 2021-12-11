package day11

import "github.com/spf13/cobra"

var Cmd = &cobra.Command{
	Use:   "day11",
	Short: "Solutions for day11",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}
