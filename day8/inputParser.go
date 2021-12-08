package day8

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type input struct {
	patterns []string
	digits   []string
}

func inputParser() <-chan input {
	inputs := make(chan input)

	go func() {
		defer close(inputs)

		f, err := os.Open("day8/input.txt")
		if err != nil {
			fmt.Println("failed to open file", err.Error())
			return
		}

		scanner := bufio.NewScanner(f)
		defer f.Close()

		for scanner.Scan() {
			line := scanner.Text()
			splitedLine := strings.Split(line, " | ")
			if len(splitedLine) != 2 {
				fmt.Println("malformed input")
				return
			}

			input := input{
				patterns: strings.Split(splitedLine[0], " "),
				digits:   strings.Split(splitedLine[1], " "),
			}

			if len(input.patterns) != 10 || len(input.digits) != 4 {
				fmt.Println("malformed input")
				return
			}

			inputs <- input
		}

	}()

	return inputs
}
