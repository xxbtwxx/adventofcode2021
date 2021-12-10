package day06

import "github.com/spf13/cobra"

var Cmd = &cobra.Command{
	Use:   "day6",
	Short: "Solutions for day 6",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	Cmd.AddCommand(&cobra.Command{
		Use:   "a",
		Short: "Solution for day 6 problem A",
		Run: func(cmd *cobra.Command, args []string) {
			task(80)
		},
	})

	Cmd.AddCommand(&cobra.Command{
		Use:   "b",
		Short: "Solution for day 6 problem B",
		Run: func(cmd *cobra.Command, args []string) {
			task(256)
		},
	})
}
