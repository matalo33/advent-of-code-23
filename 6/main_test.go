package main

import (
	"reflect"
	"testing"
)

var testData = `Time:      7  15   30
Distance:  9  40  200`

func TestParseInput(t *testing.T) {
	races := []Race{{7, 9}, {15, 40}, {30, 200}}
	result := parseInput(testData, false)

	if !reflect.DeepEqual(races, result) {
		t.Errorf("want: %v, got: %v", races, result)
	}
}

func TestParseInputKern(t *testing.T) {
	races := []Race{{71530, 940200}}
	result := parseInput(testData, true)

	if !reflect.DeepEqual(races, result) {
		t.Errorf("want: %v, got: %v", races, result)
	}
}

func TestProductSlice(t *testing.T) {
	in := []int{1, 2, 3, 4, 5, 6}
	out := productSlice(in)
	if out != 720 {
		t.Errorf("want: %v, got: %v", 720, out)
	}
}

func TestPart1(t *testing.T) {
	part1 := part1(parseInput(testData, false))
	if part1 != 288 {
		t.Errorf("want: %v, got: %v", 288, part1)
	}
}

func TestPart2(t *testing.T) {
	part1 := part1(parseInput(testData, true))
	if part1 != 71503 {
		t.Errorf("want: %v, got: %v", 71503, part1)
	}
}
