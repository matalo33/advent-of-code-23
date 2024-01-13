package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Race struct {
	time int
	distance int
}

func main() {
	file, _ := os.ReadFile("input.txt")

	fmt.Printf("Part 1: %v\n", part1(parseInput(string(file), false)))
	fmt.Printf("Part 2: %v\n", part1(parseInput(string(file), true)))
}

func parseInput(input string, kern bool) []Race {
	sections := strings.Split(input, "\n")

	if kern {
		time, _ := strconv.Atoi(strings.Replace(strings.Split(sections[0], ":")[1], " ", "", -1))
		dist, _ := strconv.Atoi(strings.Replace(strings.Split(sections[1], ":")[1], " ", "", -1))
		races := make([]Race, 1)
		races[0] = Race{time, dist}
		return races
	} else {
		times := strings.Fields(strings.Split(sections[0], ":")[1])
		distances := strings.Fields(strings.Split(sections[1], ":")[1])

		races := make([]Race, len(times))
		for i := range races {
			time, _ := strconv.Atoi(times[i])
			distance, _ := strconv.Atoi(distances[i])
			races[i] = Race{time, distance}
		}
		return races
	}
}

func part1(races []Race) int {
	var raceWins []int

	for r := range races {
		countFaster := 0
		for t := 0; t <= races[r].time; t++ {
			if (races[r].time - t) * t > races[r].distance {
				countFaster++
			}
		}
		raceWins = append(raceWins, countFaster)
	}

	return productSlice(raceWins)
}

func productSlice (s []int) int {
	r := 1
	for i := range s {
		r = r * s[i]
	}
	return r
}
