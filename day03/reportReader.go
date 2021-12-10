package day03

import (
	"bufio"
	"fmt"
	"os"
)

const reportLength = 12

func reportReader() <-chan string {
	reports := make(chan string)

	go func() {
		defer close(reports)

		f, err := os.Open("day03/input.txt")
		if err != nil {
			fmt.Println("failed to open file", err.Error())
			return
		}

		scanner := bufio.NewScanner(f)
		defer f.Close()

		for scanner.Scan() {
			report := scanner.Text()
			if len(report) != reportLength {
				fmt.Println("malformed report input")
				return
			}

			reports <- report
		}
	}()

	return reports
}
