package day09

import (
	"bufio"
	"fmt"
	"os"
)

type coords struct {
	x, y int
}

func parseMap() (floorMap [][]int) {
	f, err := os.Open("day09/input.txt")
	if err != nil {
		fmt.Println("failed to open file", err.Error())
		return
	}

	scanner := bufio.NewScanner(f)
	defer f.Close()

	for scanner.Scan() {
		line := scanner.Text()
		row := []int{}
		for _, el := range line {
			row = append(row, int(el-'0'))
		}
		floorMap = append(floorMap, row)
	}

	return
}
