package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Game struct {
	num int
	winners []int
	haves []int
	score int
}

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	games := parseInput(scanner)

	part2Games := copyGames(games)

	calculateScores(games)
	fmt.Printf("Part 1: %v\n", totalScores(games))

	fmt.Printf("Part 2: %v\n", len(part2Games))
}

func parseInput(scanner *bufio.Scanner) []Game {
	var games []Game
	num := 1

	for scanner.Scan() {
		var game Game
		game.num = num

		outer := strings.Split(scanner.Text(), "|")
		inner := strings.Split(outer[0], ":")

		for _, i := range strings.Fields(inner[1]) {
			num, _ := strconv.Atoi(string(i))
			game.winners = append(game.winners, num)
		}

		for _, i := range strings.Fields(outer[1]) {
			num, _ := strconv.Atoi(string(i))
			game.haves = append(game.haves, num)
		}

		games = append(games, game)
		num++
	}
	return games
}

func calculateScores(games []Game) {
	for idx := range games {
		e := &games[idx]
		for _, winner := range e.winners {
			if slices.Contains(e.haves, winner) {
				if e.score == 0 {
					e.score = 1
				} else {
					e.score = e.score * 2
				}
			}
		}
	}
}

func copyGames(original []Game) []Game {
	// New collection
	var games []Game
	// Loop over the original
	for idx := range original {
		// Recurse it
		recurse(&original, &games, idx)
	}
	return games
}

func recurse(original *[]Game, games *[]Game, pos int) {
	new := (*original)[pos]
	*games = append(*games, new)
	numWinners := 0
	for _, winner := range new.winners {
		if slices.Contains(new.haves, winner) {
			numWinners++
		}
	}
	for i := 0; i < numWinners; i++ {
		recurse(original, games, pos+i+1)
	}
}

func totalScores(games []Game) (total int) {
	for _, game := range games {
		total += game.score
	}
	return total
}
