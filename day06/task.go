package day06

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func task(day0s int) {
	lanternFish := [9]int{}

	f, err := os.Open("day06/input.txt")
	if err != nil {
		fmt.Println("failed to open file", err.Error())
		return
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Scan()
	input := scanner.Text()

	initialFish := strings.Split(input, ",")

	for _, fish := range initialFish {
		day0sToMultiply, err := strconv.Atoi(fish)
		if err != nil {
			fmt.Println("malformed input, value is not int")
			return
		}

		lanternFish[day0sToMultiply]++
	}

	for ; day0s != 0; day0s-- {
		newFish := 0
		for day0sToMultiply := range lanternFish {
			switch day0sToMultiply {
			case 0:
				newFish = lanternFish[day0sToMultiply]
			default:
				lanternFish[day0sToMultiply-1] = lanternFish[day0sToMultiply]
			}
		}

		lanternFish[6] += newFish
		lanternFish[8] = newFish
	}

	totalFish := 0
	for _, fishCount := range lanternFish {
		totalFish += fishCount
	}

	fmt.Println(totalFish)
}
