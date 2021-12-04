package day4

import "github.com/spf13/cobra"

var Cmd = &cobra.Command{
	Use:   "day4",
	Short: "Solutions for day4",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}
