package day2

import "github.com/spf13/cobra"

var Cmd = &cobra.Command{
	Use:   "day2",
	Short: "Solutions for day2",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}
