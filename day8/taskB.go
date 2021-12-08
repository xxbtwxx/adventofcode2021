package day8

import (
	"fmt"
	"sort"

	"github.com/spf13/cobra"
)

func init() {
	Cmd.AddCommand(&cobra.Command{
		Use:   "b",
		Short: "Solution for day 8 problem B",
		Run: func(cmd *cobra.Command, args []string) {
			taskB()
		},
	})
}

func sortString(s string) string {
	r := []byte(s)
	sort.Slice(r, func(i, j int) bool {
		return r[i] < r[j]
	})

	return string(r)
}

func isSubset(subset, set string) bool {
	setMap := map[rune]struct{}{}

	for _, el := range set {
		setMap[el] = struct{}{}
	}

	for _, el := range subset {
		if _, ok := setMap[el]; !ok {
			return false
		}
	}

	return true
}

func taskB() {
	sum := 0
	for input := range inputParser() {
		patternToValue := map[string]int{}
		valueToPattern := map[int]string{}
		undiscoveredPatternsToValue := map[string]struct{}{}

		for _, pattern := range input.patterns {
			pattern = sortString(pattern)
			switch len(pattern) {
			case 2:
				patternToValue[pattern] = 1
				valueToPattern[1] = pattern
			case 3:
				patternToValue[pattern] = 7
				valueToPattern[7] = pattern
			case 4:
				patternToValue[pattern] = 4
				valueToPattern[4] = pattern
			case 7:
				patternToValue[pattern] = 8
				valueToPattern[8] = pattern
			default:
				undiscoveredPatternsToValue[pattern] = struct{}{}
			}
		}

		// discover 0
		for pattern := range undiscoveredPatternsToValue {
			if isSubset(pattern, valueToPattern[8]) && isSubset(valueToPattern[7], pattern) && !isSubset(valueToPattern[4], pattern) && !isSubset(pattern, valueToPattern[4]) && len(pattern) == 6 {
				patternToValue[pattern] = 0
				valueToPattern[0] = pattern
				delete(undiscoveredPatternsToValue, pattern)
				break
			}
		}

		// discover 9
		for pattern := range undiscoveredPatternsToValue {
			if isSubset(valueToPattern[7], pattern) && isSubset(valueToPattern[4], pattern) && !isSubset(pattern, valueToPattern[0]) && !isSubset(valueToPattern[0], pattern) {
				patternToValue[pattern] = 9
				valueToPattern[9] = pattern
				delete(undiscoveredPatternsToValue, pattern)
				break
			}
		}

		// discover 6
		for pattern := range undiscoveredPatternsToValue {
			if len(pattern) == 6 {
				patternToValue[pattern] = 6
				valueToPattern[6] = pattern
				delete(undiscoveredPatternsToValue, pattern)
				break
			}
		}

		// discover 3
		for pattern := range undiscoveredPatternsToValue {
			if isSubset(valueToPattern[1], pattern) {
				patternToValue[pattern] = 3
				valueToPattern[3] = pattern
				delete(undiscoveredPatternsToValue, pattern)
				break
			}
		}

		// discover 5
		for pattern := range undiscoveredPatternsToValue {
			if isSubset(pattern, valueToPattern[6]) {
				patternToValue[pattern] = 5
				valueToPattern[5] = pattern
				delete(undiscoveredPatternsToValue, pattern)
				break
			}
		}

		// discover 2
		if len(undiscoveredPatternsToValue) != 1 {
			fmt.Println("unexpected case")
			return
		}

		for pattern := range undiscoveredPatternsToValue {
			patternToValue[pattern] = 2
			valueToPattern[2] = pattern
			delete(undiscoveredPatternsToValue, pattern)
		}

		mutliplier := 1000
		for _, signal := range input.signal {
			sum += patternToValue[sortString(signal)] * mutliplier
			mutliplier /= 10
		}
	}

	fmt.Println(sum)
}
