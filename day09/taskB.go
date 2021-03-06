package day09

import (
	"fmt"
	"sort"

	"github.com/spf13/cobra"
)

func init() {
	Cmd.AddCommand(&cobra.Command{
		Use:   "b",
		Short: "Solution for day 9 problem B",
		Run: func(cmd *cobra.Command, args []string) {
			taskB()
		},
	})
}

func taskB() {
	floorMap := parseMap()
	lowCoords := []coords{}

	for y := 0; y < len(floorMap); y++ {
		for x := 0; x < len(floorMap[y]); x++ {
			sorted := getSortedAdjecent(floorMap, x, y)
			if floorMap[y][x] < sorted[0] {
				lowCoords = append(lowCoords, coords{x: x, y: y})
			}
		}
	}

	iteratedOver := map[coords]struct{}{}
	regionsSize := []int{}
	for _, coord := range lowCoords {
		regionSize := 1
		iteratedOver[coord] = struct{}{}
		adjecent := adjecentCoords(floorMap, iteratedOver, coord.x, coord.y)
		for len(adjecent) != 0 {
			adjecentCoord := adjecent[0]
			adjecent = adjecent[1:]
			if _, ok := iteratedOver[adjecentCoord]; !ok {
				iteratedOver[adjecentCoord] = struct{}{}
				regionSize++
				adjecent = append(adjecent, adjecentCoords(floorMap, iteratedOver, adjecentCoord.x, adjecentCoord.y)...)
			}
		}

		regionsSize = append(regionsSize, regionSize)
	}

	sort.Ints(regionsSize)
	regionsMult := regionsSize[len(regionsSize)-1] * regionsSize[len(regionsSize)-2] * regionsSize[len(regionsSize)-3]
	fmt.Println(regionsMult)
}

func adjecentCoords(m [][]int, iterated map[coords]struct{}, x, y int) []coords {
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
		if _, ok := iterated[pair]; !ok && m[pair.y][pair.x] != 9 {
			adjecent = append(adjecent, pair)
		}
	}

	return adjecent
}
