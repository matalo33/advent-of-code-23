package main

import "testing"

var testData = struct {
	input string
}{
	`467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`,
}

func Test_parseInput(t *testing.T) {
	schematic := parseInput(testData.input)
	if schematic[0][0].num != 4 {
		t.Errorf("0.0 should be 4: %v", schematic[0][0].num)
	}

	if schematic[9][9].isNum != false || schematic[9][9].isSymbol != false {
		t.Errorf("9.9 should be not num and not symbol: %v %v", schematic[9][9].isNum, schematic[9][9].isSymbol)
	}

	if !schematic[1][3].isSymbol {
		t.Errorf("1.3 should be a symbol: %v", schematic[1][3].isSymbol)
	}
}

func Test_concatNumber(t *testing.T) {
	result := concatNumbers([]int{1, 2, 4, 9, 2})
	if result != 12492 {
		t.Errorf("want: %v, got: %v", 12492, result)
	}
}

func Test_part1(t *testing.T) {
	schematic := parseInput(testData.input)
	result := part1(schematic)
	if result != 4361 {
		t.Errorf("want: %v, got: %v", 4361, result)
	}
}
