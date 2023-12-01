package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.ReadFile("input.txt")
	input := string(file)

	fmt.Printf("Part 1: %v\n", solve(input))

	input = strings.ReplaceAll(input, "one", "o1e")
	input = strings.ReplaceAll(input, "two", "t2e")
	input = strings.ReplaceAll(input, "three", "thr3e")
	input = strings.ReplaceAll(input, "four", "fo4r")
	input = strings.ReplaceAll(input, "five", "fi5e")
	input = strings.ReplaceAll(input, "six", "s6x")
	input = strings.ReplaceAll(input, "seven", "se7en")
	input = strings.ReplaceAll(input, "eight", "ei8ht")
	input = strings.ReplaceAll(input, "nine", "ni9e")

	fmt.Printf("Part 2: %v\n", solve(input))
}

func solve(input string) int {
	var values []int
	for _, line := range strings.Split(input, "\n") {
		var first, last int
		var err error
		for _, c := range line {
			first, err = strconv.Atoi(string(c))
			if err == nil {
				break
			}
		}
		err = nil
		for i := len(line) - 1; i >= 0; i-- {
			last, err = strconv.Atoi(string(line[i]))
			if err == nil {
				break
			}
		}

		//fmt.Printf("%v: %v %v\n", line, first, last)
		value, _ := strconv.Atoi(fmt.Sprintf("%v%v", first, last))
		values = append(values, value)
	}

	total := 0
	for _, a := range values {
		total += a
	}
	return total
}
