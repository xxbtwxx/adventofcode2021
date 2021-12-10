package day10

import (
	"bufio"
	"fmt"
	"os"
)

func inputReader() <-chan string {
	inputs := make(chan string)

	go func() {
		defer close(inputs)

		f, err := os.Open("day10/input.txt")
		if err != nil {
			fmt.Println("failed to open file", err.Error())
			return
		}

		scanner := bufio.NewScanner(f)
		defer f.Close()

		for scanner.Scan() {
			inputs <- scanner.Text()
		}

	}()

	return inputs
}
