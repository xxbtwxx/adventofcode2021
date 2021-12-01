package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	depthValues := depthReader()

	elements := make([]int, 0, 4)
	increments := 0
	for depthValue := range depthValues {
		elements = append(elements, depthValue)
		if len(elements) < 4 {
			continue
		}

		if sum(elements[:len(elements)-1]) < sum(elements[1:]) {
			increments++
		}

		elements = elements[1:]
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

func sum(elements []int) (s int) {
	for _, el := range elements {
		s += el
	}

	return
}
