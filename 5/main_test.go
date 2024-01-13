package main

import (
	"testing"
)

var testData = `seeds: 79 14 55 13

seed-to-soil map:
50 98 2
52 50 48

soil-to-fertilizer map:
0 15 37
37 52 2
39 0 15

fertilizer-to-water map:
49 53 8
0 11 42
42 0 7
57 7 4

water-to-light map:
88 18 7
18 25 70

light-to-temperature map:
45 77 23
81 45 19
68 64 13

temperature-to-humidity map:
0 69 1
1 0 69

humidity-to-location map:
60 56 37
56 93 4`

func Test_parseInput(t *testing.T) {
	seeds, maps := parseInput(testData)
	if len(seeds) != 4 {
		t.Errorf("Seed length wrong: %v", len(seeds))
	}
	if len(maps["fertilizer-to-water"]) != 4 {
		t.Errorf("fertilizer-to-water length wrong:, %v", len(maps["fertilizer-to-water"]))
	}
	if maps["temperature-to-humidity"][1].length != 69 {
		t.Errorf("temperature-to-humidity 2 length is wrong: %v", maps["temperature-to-humidity"][1].length)
	}
	if maps["seed-to-soil"][0].dest != 50 {
		t.Errorf("seed-to-soil 1 dest should be 50: %v", maps["seed-to-soil"][0].dest)
	}
}

func Test_part1(t *testing.T) {
	seeds, maps := parseInput(testData)
	ans := part1(seeds, maps)
	if ans != 35 {
		t.Errorf("Part 1 want: %v, got: %v", 35, ans)
	}
}

func Test_SeedsToSeeds(t *testing.T) {
	seeds, _ := parseInput(testData)
	seeds = seedToSeeds(seeds)
	if len(seeds) != 27 {
		t.Errorf("seeds to seeds not right should be 27: %v", len(seeds))
	}
}

func Test_part2(t *testing.T) {
	seeds, maps := parseInput(testData)
	seeds = seedToSeeds(seeds)
	ans := part1(seeds, maps)
	if ans != 46 {
		t.Errorf("Part 2 want: %v, got: %v", 46, ans)
	}
}
