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
		discoveredPatterns := map[string]int{}
		valueToPattern := map[int]string{}
		undiscoveredPatterns := map[string]struct{}{}

		for _, pattern := range input.patterns {
			undiscoveredPatterns[sortString(pattern)] = struct{}{}
		}

		// first pass, discover unique digits
		for pattern := range undiscoveredPatterns {
			switch len(pattern) {
			case 2:
				discoveredPatterns[pattern] = 1
				valueToPattern[1] = pattern
				delete(undiscoveredPatterns, pattern)
			case 3:
				discoveredPatterns[pattern] = 7
				valueToPattern[7] = pattern
				delete(undiscoveredPatterns, pattern)
			case 4:
				discoveredPatterns[pattern] = 4
				valueToPattern[4] = pattern
				delete(undiscoveredPatterns, pattern)
			case 7:
				discoveredPatterns[pattern] = 8
				valueToPattern[8] = pattern
				delete(undiscoveredPatterns, pattern)
			}
		}

		// discover 0
		for pattern := range undiscoveredPatterns {
			if isSubset(pattern, valueToPattern[8]) && isSubset(valueToPattern[7], pattern) && !isSubset(valueToPattern[4], pattern) && !isSubset(pattern, valueToPattern[4]) && len(pattern) == 6 {
				discoveredPatterns[pattern] = 0
				valueToPattern[0] = pattern
				delete(undiscoveredPatterns, pattern)
				break
			}
		}

		// discover 9
		for pattern := range undiscoveredPatterns {
			if isSubset(valueToPattern[7], pattern) && isSubset(valueToPattern[4], pattern) && !isSubset(pattern, valueToPattern[0]) && !isSubset(valueToPattern[0], pattern) {
				discoveredPatterns[pattern] = 9
				valueToPattern[9] = pattern
				delete(undiscoveredPatterns, pattern)
				break
			}
		}

		// discover 6
		for pattern := range undiscoveredPatterns {
			if len(pattern) == 6 {
				discoveredPatterns[pattern] = 6
				valueToPattern[6] = pattern
				delete(undiscoveredPatterns, pattern)
				break
			}
		}

		// discover 3
		for pattern := range undiscoveredPatterns {
			if isSubset(valueToPattern[1], pattern) {
				discoveredPatterns[pattern] = 3
				valueToPattern[3] = pattern
				delete(undiscoveredPatterns, pattern)
				break
			}
		}

		// discover 5
		for pattern := range undiscoveredPatterns {
			if isSubset(pattern, valueToPattern[6]) {
				discoveredPatterns[pattern] = 5
				valueToPattern[5] = pattern
				delete(undiscoveredPatterns, pattern)
				break
			}
		}

		// discover 2
		if len(undiscoveredPatterns) != 1 {
			fmt.Println("unexpected case")
			return
		}

		for pattern := range undiscoveredPatterns {
			discoveredPatterns[pattern] = 2
			valueToPattern[2] = pattern
			delete(undiscoveredPatterns, pattern)
		}

		mutliplier := 1000
		for _, signal := range input.signal {
			sum += discoveredPatterns[sortString(signal)] * mutliplier
			mutliplier /= 10
		}
	}

	fmt.Println(sum)
}
