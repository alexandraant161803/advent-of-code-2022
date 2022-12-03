package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

/*

Scoring for choice:
- 1 for Rock
- 2 for Paper
- 3 for Scissors

Scoring for outcome of the round:
- 0 for a loss
- 3 for a draw
- 6 for a win

*/

func main() {
	// Read the file
	input, err := os.ReadFile("day2input.txt")
	if err != nil {
		log.Fatal(err)
	}

	games := strings.Split(string(input), "\n")
	sumPart1 := 0
	sumPart2 := 0
	for _, game := range games {
		round := strings.Split(string(game), " ")
		sumPart1 += part1(round)
		sumPart2 += part2(round)
	}
	fmt.Println("Part 1: ", sumPart1)
	fmt.Println("Part 2: ", sumPart2)
}

/*

Part 1:

For opponent:
- Rock = A
- Paper = B
- Scissors = C

For you:
- Rock = X
- Paper = Y
- Scissors = Z

*/

/*What would your total score be if everything goes exactly according to your strategy guide?*/
func part1(round []string) int {
	myPoints := 0
	them := round[0]
	you := round[1]

	switch {
	case them == "A":
		if you == "X" {
			// Tie (3) and rock (1)
			myPoints += 4

		} else if you == "Y" {
			// Win (6) and paper (2)
			myPoints += 8

		} else {
			// Lose (0) and scissors (3)
			myPoints += 3

		}
	case them == "B":
		if you == "X" {
			// loss (0) and rock (1)
			myPoints += 1
		} else if you == "Y" {
			// Tie (3) and paper (2)
			myPoints += 5

		} else {
			// Win (6) and scissors (3)
			myPoints += 9
		}
	case them == "C":
		if you == "X" {
			// Win (6) and rock (1)
			myPoints += 7
		} else if you == "Y" {
			// Lose (0) and paper (2)
			myPoints += 2
		} else {
			// Tie (3) and scissors (3)
			myPoints += 6
		}

	}

	return myPoints

}

/*
For opponent:
- Rock = A
- Paper = B
- Scissors = C

For you:
- X = need to lose
- Y = need to tie
- Z = need to win

*/

/*What would your total score be if everything goes exactly according to your strategy guide?*/
func part2(round []string) int {
	myPoints := 0
	them := round[0]
	you := round[1]

	switch {
	case them == "A":
		if you == "X" {
			// Need to lose (0) and scissors (3)
			myPoints += 3

		} else if you == "Y" {
			// Need to tie (3) and rock (1)
			myPoints += 4

		} else {
			// Need to win (6) and paper (2)
			myPoints += 8

		}

	case them == "B":
		if you == "X" {
			// Need to lose (0) and rock (1)
			myPoints += 1
		} else if you == "Y" {
			// Need to tie (3) and paper (2)
			myPoints += 5

		} else {
			// Need to win (6) and scissors (3)
			myPoints += 9
		}
	case them == "C":
		if you == "X" {
			// Need to lose (0) and paper (2)
			myPoints += 2
		} else if you == "Y" {
			// Need to tie (3) and scissors (3)
			myPoints += 6
		} else {
			// Need to win (6) and rock (1)
			myPoints += 7
		}

	}

	return myPoints
}
