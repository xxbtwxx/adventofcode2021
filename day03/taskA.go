package day03

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	Cmd.AddCommand(&cobra.Command{
		Use:   "a",
		Short: "Solution for day 3 problem A",
		Run: func(cmd *cobra.Command, args []string) {
			taskA()
		},
	})
}

func taskA() {
	reportsCount := 0
	onesCounter := [reportLength]int{}
	for report := range reportReader() {
		reportsCount++
		for index, value := range report {
			onesCounter[index] += runeToVal[value]
		}
	}

	gammaValues, epsilonValues := [reportLength]int{}, [reportLength]int{}
	for index, val := range onesCounter {
		if val > reportsCount/2 {
			gammaValues[index] = 1
		} else {
			epsilonValues[index] = 1
		}
	}

	gammaValue := binaryIntSliceToInt(gammaValues[:])
	epsilonValue := binaryIntSliceToInt(epsilonValues[:])

	fmt.Println(gammaValue * epsilonValue)
}
