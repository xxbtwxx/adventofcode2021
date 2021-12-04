package day4

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func inputReader() (<-chan int, map[int]*bingoCard) {
	drawnNum := make(chan int)
	cards := map[int]*bingoCard{}

	f, err := os.Open("day4/input.txt")
	if err != nil {
		fmt.Println("failed to open file", err.Error())
		return drawnNum, cards
	}

	scanner := bufio.NewScanner(f)
	defer f.Close()

	scanner.Scan()
	draws := scanner.Text()
	drawnNums := strings.Split(draws, ",")

	go func() {
		defer close(drawnNum)
		for _, number := range drawnNums {
			num, err := strconv.Atoi(number)
			if err != nil {
				fmt.Println("value is not a number")
				return
			}

			drawnNum <- num
		}
	}()

	scanner.Scan()

	cardValues := [5][5]int{}
	row := 0
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		line = strings.ReplaceAll(line, "  ", " ")
		line = strings.Trim(line, " ")

		rowValues := strings.Split(line, " ")
		rowVals := [5]int{}
		for index, value := range rowValues {
			numValue, err := strconv.Atoi(value)
			if err != nil {
				fmt.Println("bingo card value is not a number")
				return drawnNum, cards
			}

			rowVals[index] = numValue
		}

		cardValues[row] = rowVals

		row++
		if row == 5 {
			cards[len(cards)+1] = initializeCard(cardValues)
			row = 0
		}
	}

	return drawnNum, cards
}
