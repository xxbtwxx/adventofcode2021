package day6

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func task(days int) {
	lanternFish := [9]int{}

	f, err := os.Open("day6/input.txt")
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
		daysToMultiply, err := strconv.Atoi(fish)
		if err != nil {
			fmt.Println("malformed input, value is not int")
			return
		}

		lanternFish[daysToMultiply]++
	}

	for ; days != 0; days-- {
		newFish := 0
		for daysToMultiply := range lanternFish {
			switch daysToMultiply {
			case 0:
				newFish = lanternFish[daysToMultiply]
			default:
				lanternFish[daysToMultiply-1] = lanternFish[daysToMultiply]
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
