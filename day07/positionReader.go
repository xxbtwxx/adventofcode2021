package day07

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func positionReader() <-chan int {
	positions := make(chan int)

	go func() {
		defer close(positions)

		f, err := os.Open("day07/input.txt")
		if err != nil {
			fmt.Println("failed to open file", err.Error())
			return
		}

		scanner := bufio.NewScanner(f)
		defer f.Close()
		scanner.Scan()

		positionsSlice := strings.Split(scanner.Text(), ",")
		for _, position := range positionsSlice {
			pos, err := strconv.Atoi(position)
			if err != nil {
				fmt.Println("malformed input, position is not int")
				return
			}

			positions <- pos
		}

	}()

	return positions
}
