package day9

import "github.com/spf13/cobra"

var Cmd = &cobra.Command{
	Use:   "day9",
	Short: "Solutions for day9",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}
