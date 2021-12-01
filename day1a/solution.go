package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	depthValues := depthReader()

	previousValue := <-depthValues
	increments := 0
	for depthValue := range depthValues {
		if depthValue > previousValue {
			increments++
		}

		previousValue = depthValue
	}

	fmt.Println(increments)
}

func depthReader() <-chan int {
	depthValues := make(chan int)

	go func() {
		defer close(depthValues)

		f, err := os.Open("input.txt")
		if err != nil {
			fmt.Println("failed to open file", err.Error())
			return
		}

		scanner := bufio.NewScanner(f)
		defer f.Close()

		for scanner.Scan() {
			depth, err := strconv.Atoi(scanner.Text())
			if err != nil {
				fmt.Println("failed string to int conversion", err.Error())
				return
			}
			depthValues <- depth
		}
	}()

	return depthValues
}
