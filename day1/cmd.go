package day1

import "github.com/spf13/cobra"

var Cmd = &cobra.Command{
	Use:   "day1",
	Short: "Solutions for day1",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}
