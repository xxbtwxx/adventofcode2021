package day05

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func lineReader() <-chan line {
	lines := make(chan line)

	go func() {
		defer close(lines)

		f, err := os.Open("day05/input.txt")
		if err != nil {
			fmt.Println("failed to open file", err.Error())
			return
		}

		scanner := bufio.NewScanner(f)
		defer f.Close()

		for scanner.Scan() {
			inputLine := scanner.Text()
			linePoints := strings.Split(inputLine, " -> ")
			if len(linePoints) != 2 {
				fmt.Println("malformed line input")
				return
			}

			startPoint := strings.Split(linePoints[0], ",")
			endPoint := strings.Split(linePoints[1], ",")

			startX, err := strconv.Atoi(startPoint[0])
			if err != nil {
				fmt.Println("startX point is not an int")
				return
			}

			startY, err := strconv.Atoi(startPoint[1])
			if err != nil {
				fmt.Println("startY point is not an int")
				return
			}

			endX, err := strconv.Atoi(endPoint[0])
			if err != nil {
				fmt.Println("startY point is not an int")
				return
			}

			endY, err := strconv.Atoi(endPoint[1])
			if err != nil {
				fmt.Println("startY point is not an int")
				return
			}

			lines <- line{start: point{x: startX, y: startY}, end: point{x: endX, y: endY}}
		}
	}()

	return lines

}
