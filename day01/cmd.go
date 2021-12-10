package day01

import "github.com/spf13/cobra"

var Cmd = &cobra.Command{
	Use:   "day1",
	Short: "Solutions for day 1",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	Cmd.AddCommand(&cobra.Command{
		Use:   "a",
		Short: "Solution for day 1 problem A",
		Run: func(cmd *cobra.Command, args []string) {
			task(1)
		},
	})

	Cmd.AddCommand(&cobra.Command{
		Use:   "b",
		Short: "Solution for day 1 problem B",
		Run: func(cmd *cobra.Command, args []string) {
			task(3)
		},
	})
}
