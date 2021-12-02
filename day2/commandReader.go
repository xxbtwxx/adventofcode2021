package day2

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type command struct {
	direction string
	value     int
}

func commandReader() <-chan command {
	commands := make(chan command)

	go func() {
		defer close(commands)

		f, err := os.Open("day2/input.txt")
		if err != nil {
			fmt.Println("failed to open file", err.Error())
			return
		}

		scanner := bufio.NewScanner(f)
		defer f.Close()

		for scanner.Scan() {
			line := scanner.Text()
			splittedCommand := strings.Split(line, " ")
			if len(splittedCommand) != 2 {
				fmt.Println("malformed input")
				return
			}

			value, err := strconv.Atoi(splittedCommand[1])
			if err != nil {
				fmt.Println("failed string to int conversion", err.Error())
				return
			}

			command := command{
				direction: splittedCommand[0],
				value:     value,
			}

			commands <- command
		}
	}()

	return commands
}
