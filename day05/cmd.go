package day05

import "github.com/spf13/cobra"

var Cmd = &cobra.Command{
	Use:   "day5",
	Short: "Solutions for day 5",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	Cmd.AddCommand(&cobra.Command{
		Use:   "a",
		Short: "Solution for day 5 problem A",
		Run: func(cmd *cobra.Command, args []string) {
			task(false)
		},
	})

	Cmd.AddCommand(&cobra.Command{
		Use:   "b",
		Short: "Solution for day 5 problem B",
		Run: func(cmd *cobra.Command, args []string) {
			task(true)
		},
	})
}
