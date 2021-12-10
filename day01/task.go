package day01

import (
	"fmt"
)

func task(sumSize int) {
	elements := make([]int, 0, sumSize+1)
	increments := 0
	for depthValue := range depthReader() {
		elements = append(elements, depthValue)
		if len(elements) < sumSize+1 {
			continue
		}

		if elements[0] < elements[len(elements)-1] {
			increments++
		}

		elements = elements[1:]
	}

	fmt.Println(increments)
}
