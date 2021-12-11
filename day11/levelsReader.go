package day11

import (
	"bufio"
	"fmt"
	"os"
)

func energyLevelsReader() [][]int {
	energyLevels := [][]int{}

	f, err := os.Open("day11/input.txt")
	if err != nil {
		fmt.Println("failed to open file", err.Error())
		return energyLevels
	}

	scanner := bufio.NewScanner(f)
	defer f.Close()

	for scanner.Scan() {
		line := scanner.Text()
		energyLevel := []int{}

		for _, val := range line {
			energyLevel = append(energyLevel, int(val-'0'))
		}

		energyLevels = append(energyLevels, energyLevel)
	}
	return energyLevels
}
