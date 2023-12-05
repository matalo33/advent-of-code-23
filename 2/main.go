package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Game struct {
	red   int
	blue  int
	green int
}

func main() {
	file, _ := os.ReadFile(("input.txt"))

	games := parseInput(string(file))

	fmt.Printf("Part 1: %v\n", part1(games, 12, 13, 14))
	fmt.Printf("Part 2: %v\n", part2(games, 12, 13, 14))
}

func parseInput(input string) [][]Game {
	lines := strings.Split(input, "\n")
	games := make([][]Game, len(lines)-1)
	for i := 0; i < len(lines)-1; i++ {
		var thisRound []Game
		gameInput := strings.Split(lines[i], ":")
		gameRounds := strings.Split(gameInput[1], ";")
		for _, gameRound := range gameRounds {
			draws := strings.Split(gameRound, ",")
			var newGame Game
			for _, draw := range draws {
				drawParts := strings.Split(strings.TrimSpace(draw), " ")
				num, _ := strconv.Atoi(drawParts[0])
				if drawParts[1] == "red" {
					newGame.red = num
				} else if drawParts[1] == "green" {
					newGame.green = num
				} else if drawParts[1] == "blue" {
					newGame.blue = num
				}
			}
			thisRound = append(thisRound, newGame)
		}
		games[i] = thisRound
	}
	return games
}

func part1(games [][]Game, maxRed int, maxGreen int, maxBlue int) (result int) {
	for i, game := range games {
		passing := true
		for _, subgame := range game {
			if subgame.red > maxRed || subgame.green > maxGreen || subgame.blue > maxBlue {
				passing = false
			}
		}
		if passing {
			result = result + (i + 1)
		}
	}
	return result
}

func part2(games [][]Game, maxRed int, maxGreen int, maxBlue int) (result int) {
	for _, game := range games {
		var minRed, minGreen, minBlue int
		for _, subgame := range game {
			minRed = max(minRed, subgame.red)
			minGreen = max(minGreen, subgame.green)
			minBlue = max(minBlue, subgame.blue)
		}
		result = result + (minRed * minGreen * minBlue)
	}
	return result
}
