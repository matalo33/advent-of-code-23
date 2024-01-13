package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Map struct {
	source int
	dest int
	length int
}

var order = []string{
	"seed-to-soil",
	"soil-to-fertilizer",
	"fertilizer-to-water",
	"water-to-light",
	"light-to-temperature",
	"temperature-to-humidity",
	"humidity-to-location",
}

func main() {
	file, _ := os.ReadFile("input.txt")

	seeds, maps := parseInput(string(file))

	fmt.Printf("Part 1: %v\n", part1(seeds, maps))
	newSeeds := seedToSeeds(seeds)
	fmt.Printf("Part 2: %v\n", part1(newSeeds, maps))
}

func parseInput(input string) ([]int, map[string][]Map) {
	sections := strings.Split(input, "\n\n")
	var seeds []int
	maps := make(map[string][]Map)

	for _, s := range strings.Fields(strings.Split(sections[0], ":")[1]) {
		i, _ := strconv.Atoi(s)
		seeds = append(seeds, i)
	}

	for i, o := range order {
		var sm []Map
		for _, m := range strings.Split(strings.Split(sections[i+1], ":\n")[1], "\n") {
			n := strings.Fields(m)
			d, _ := strconv.Atoi(n[0])
			s, _ := strconv.Atoi(n[1])
			l, _ := strconv.Atoi(n[2])

			sm = append(sm, Map{
				source: s,
				dest: d,
				length: l,
			})
		}
		maps[o] = sm
	}

	return seeds, maps
}

func part1(seeds []int, maps map[string][]Map) int {
	var answers []int
	fmt.Printf("There are %v seeds\n", len(seeds))
	for i, seed := range seeds {
		// For each type(order) of map
		if i % 10000000 == 0 {
			fmt.Printf("Remaining: %v\n", (len(seeds)-i))
		}
		ans := seed
		for _, o := range order {
			// For each map
			for _, m := range maps[o] {
				// Can we find the seed within range?
				if ans >= m.source && ans <= (m.source + m.length) {
					ans = m.dest + (ans - m.source)
					break
				}
			}
		}
		answers = append(answers, ans)
	}
	return slices.Min(answers)
}

func seedToSeeds(seeds []int) []int {
	var newSeeds []int
	for i := 0; i < len(seeds); i+=2 {
		for s := 0; s < seeds[i+1]; s++ {
			newSeeds = append(newSeeds, seeds[i]+s)
		}
	}
	return newSeeds
}
