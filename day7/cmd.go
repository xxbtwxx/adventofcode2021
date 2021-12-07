package day7

import "github.com/spf13/cobra"

var Cmd = &cobra.Command{
	Use:   "day7",
	Short: "Solutions for day7",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}
