package day9

import (
	"fmt"
	"sort"

	"github.com/spf13/cobra"
)

func init() {
	Cmd.AddCommand(&cobra.Command{
		Use:   "a",
		Short: "Solution for day 9 problem A",
		Run: func(cmd *cobra.Command, args []string) {
			taskA()
		},
	})
}

func taskA() {
	lowsSum := 0
	floorMap := parseMap()

	for y := 0; y < len(floorMap); y++ {
		for x := 0; x < len(floorMap[y]); x++ {
			sorted := getSortedAdjecent(floorMap, x, y)
			if floorMap[y][x] < sorted[0] {
				lowsSum += floorMap[y][x] + 1
			}
		}
	}

	fmt.Println(lowsSum)
}

func getSortedAdjecent(m [][]int, x, y int) []int {
	adjecent := []int{}

	adjecentPairs := map[coords]struct{}{
		{x: x + 1, y: y}: {},
		{x: x - 1, y: y}: {},
		{x: x, y: y + 1}: {},
		{x: x, y: y - 1}: {},
	}

	for coord := range adjecentPairs {
		if coord.x < 0 || coord.x >= len(m[y]) || coord.y < 0 || coord.y >= len(m) {
			delete(adjecentPairs, coord)
		}
	}

	for coord := range adjecentPairs {
		adjecent = append(adjecent, m[coord.y][coord.x])
	}

	sort.Ints(adjecent)
	return adjecent
}

func adjecentCoords(m [][]int, x, y int) []coords {
	adjecent := []coords{}

	adjecentPairs := map[coords]struct{}{
		{x: x + 1, y: y}: {},
		{x: x - 1, y: y}: {},
		{x: x, y: y + 1}: {},
		{x: x, y: y - 1}: {},
	}

	for coord := range adjecentPairs {
		if coord.x < 0 || coord.x >= len(m[y]) || coord.y < 0 || coord.y >= len(m) {
			delete(adjecentPairs, coord)
		}
	}

	for pair := range adjecentPairs {
		adjecent = append(adjecent, pair)
	}

	return adjecent
}
