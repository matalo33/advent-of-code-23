package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Bit struct {
	isSymbol bool
	isNum    bool
	num      int
	symbol   rune
}

func main() {
	file, _ := os.ReadFile(("input.txt"))

	schematic := parseInput(string(file))

	fmt.Printf("Part 1: %v\n", part1(schematic))
	fmt.Printf("Part 2: %v\n", part2(schematic, string(file)))
}

func parseInput(input string) [][]Bit {
	lines := strings.Split(input, "\n")
	gridSize := len(lines)

	schematic := make([][]Bit, gridSize)
	for l := 0; l < len(lines); l++ {
		scLine := make([]Bit, len(lines))
		for i, b := range lines[l] {
			var bit Bit
			if b != '.' {
				n, err := strconv.Atoi(string(b))
				if err != nil {
					bit.isSymbol = true
					bit.symbol = b
				} else {
					bit.isNum = true
					bit.num = n
				}
			}
			scLine[i] = bit
		}
		schematic[l] = scLine
	}
	return schematic
}

func part1(schematic [][]Bit) (result int) {
	for i, line := range schematic {
		p := 0
		for p < len(line) {
			if line[p].isNum {
				lengthOfNumber := 1
				digits := []int{line[p].num}
				for {
					if line[p+lengthOfNumber].isNum {
						digits = append(digits, line[p+lengthOfNumber].num)
						lengthOfNumber++
					} else {
						break
					}
				}
				num := concatNumbers(digits)
				numStartsAt := p
				numEndsAt := p + lengthOfNumber //Is not a number

				foundSymbol := false
				// Above
				if i > 0 {
					for l := numStartsAt - 1; l <= numEndsAt; l++ {
						if l >= 0 && l < len(line) {
							if schematic[i-1][l].isSymbol {
								foundSymbol = true
								break
							}
						}
					}
				}
				// Below
				if i+1 < len(line) {
					for l := numStartsAt - 1; l <= numEndsAt; l++ {
						if l >= 0 && l < len(line) {
							if schematic[i+1][l].isSymbol {
								foundSymbol = true
								break
							}
						}
					}
				}
				// Side to side
				if numEndsAt < len(line) {
					if schematic[i][numEndsAt].isSymbol {
						foundSymbol = true
					}
				}
				if numStartsAt > 0 {
					if schematic[i][numStartsAt-1].isSymbol {
						foundSymbol = true
					}
				}
				if foundSymbol {
					result += num
				}
				p += lengthOfNumber
			} else {
				p++
			}
		}
	}
	return
}

// part2 is a disaster. I should have regex parsed the input to start with
// so here I barely use the [][]Bit collection because I re-parse the input
// and wasn't bothered to go back and rewrite part1
func part2(schematic [][]Bit, input string) (result int) {
	lines := strings.Split(input, "\n")
	partIndexes := make([][][]int, len(lines))
	re := regexp.MustCompile(`[0-9]+`)
	for i, line := range lines {
		partIndexes[i] = re.FindAllStringIndex(line, -1)
	}

	for i, line := range schematic {
		for p := 0; p < len(line); p++ {
			if line[p].symbol == '*' {
				var adjacentSymbols []int
				// Above
				for _, partIndex := range partIndexes[i-1] {
					if (p >= partIndex[0] && p <= partIndex[1]) || (p+1 >= partIndex[0] && p+1 <= partIndex[1]) {
						adjacentSymbols = append(adjacentSymbols, extractPartNum(schematic[i-1], partIndex))
					}
				}
				// Below
				for _, partIndex := range partIndexes[i+1] {
					if (p >= partIndex[0] && p <= partIndex[1]) || (p+1 >= partIndex[0] && p+1 <= partIndex[1]) {
						adjacentSymbols = append(adjacentSymbols, extractPartNum(schematic[i+1], partIndex))
					}
				}
				// Side to side
				for _, partIndex := range partIndexes[i] {
					if (p >= partIndex[0] && p <= partIndex[1]) || (p+1 >= partIndex[0] && p+1 <= partIndex[1]) {
						adjacentSymbols = append(adjacentSymbols, extractPartNum(schematic[i], partIndex))
					}
				}
				if len(adjacentSymbols) == 2 {
					result += (adjacentSymbols[0] * adjacentSymbols[1])
				}
			}
		}
	}
	return
}

func concatNumbers(n []int) (result int) {
	var str string
	for _, i := range n {
		str = str + strconv.FormatInt(int64(i), 10)
	}
	result, _ = strconv.Atoi(str)
	return
}

func extractPartNum(line []Bit, location []int) int {
	var set []int
	for i := location[0]; i < location[1]; i++ {
		set = append(set, line[i].num)
	}
	return concatNumbers(set)
}
