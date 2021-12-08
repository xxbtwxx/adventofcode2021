package day8

import "github.com/spf13/cobra"

var Cmd = &cobra.Command{
	Use:   "day8",
	Short: "Solutions for day8",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}
