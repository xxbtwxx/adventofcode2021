package day3

import (
	"fmt"
	"sort"

	"github.com/spf13/cobra"
)

func init() {
	Cmd.AddCommand(&cobra.Command{
		Use:   "b",
		Short: "Solution for day 3 problem B",
		Run: func(cmd *cobra.Command, args []string) {
			taskB()
		},
	})
}

func taskB() {
	reports := []string{}
	for report := range reportReader() {
		reports = append(reports, report)
	}

	sort.Strings(reports)

	o2Reading := calculateO2(reports, 0)
	o2Level := binaryStringToInt(o2Reading)

	co2Reading := calculateCO2(reports, 0)
	co2Level := binaryStringToInt(co2Reading)

	fmt.Println(co2Level * o2Level)
}

func calculateO2(reports []string, index int) string {
	if index == reportLength {
		return reports[0]
	}

	zeroesCount := 0
	for _, el := range reports {
		if el[index] == '0' {
			zeroesCount++
		}
	}

	if zeroesCount > len(reports)-zeroesCount {
		if len(reports[:zeroesCount]) == 0 {
			return calculateO2(reports[zeroesCount:], index+1)
		}
		return calculateO2(reports[:zeroesCount], index+1)
	} else {
		if len(reports[zeroesCount:]) == 0 {
			return calculateO2(reports[:zeroesCount], index+1)
		}
		return calculateO2(reports[zeroesCount:], index+1)
	}
}

func calculateCO2(reports []string, index int) string {
	if index == reportLength {
		return reports[0]
	}

	zeroesCount := 0
	for _, el := range reports {
		if el[index] == '0' {
			zeroesCount++
		}
	}

	if zeroesCount > len(reports)-zeroesCount {
		if len(reports[zeroesCount:]) == 0 {
			return calculateCO2(reports[:zeroesCount], index+1)
		}
		return calculateCO2(reports[zeroesCount:], index+1)
	} else {
		if len(reports[:zeroesCount]) == 0 {
			return calculateCO2(reports[zeroesCount:], index+1)
		}
		return calculateCO2(reports[:zeroesCount], index+1)
	}
}
