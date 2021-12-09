package day9

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
		adjecent := adjecentCoords(floorMap, coord.x, coord.y)
		for len(adjecent) != 0 {
			adjecentCoord := adjecent[0]
			adjecent = adjecent[1:]
			if _, ok := iteratedOver[adjecentCoord]; !ok && floorMap[adjecentCoord.y][adjecentCoord.x] != 9 {
				iteratedOver[adjecentCoord] = struct{}{}
				regionSize++
				adjecent = append(adjecent, adjecentCoords(floorMap, adjecentCoord.x, adjecentCoord.y)...)
			}
		}

		regionsSize = append(regionsSize, regionSize)
	}

	sort.Ints(regionsSize)
	regionsMult := regionsSize[len(regionsSize)-1] * regionsSize[len(regionsSize)-2] * regionsSize[len(regionsSize)-3]
	fmt.Println(regionsMult)
}
