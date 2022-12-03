package main

import (
	"bufio"
	"fmt"
	"os"
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

Part 1:

For opponent:
- Rock = A
- Paper = B
- Scissors = C

For you:
- Rock = X
- Paper = Y
- Scissors = Z

Part 2:
For opponent:
- Rock = A
- Paper = B
- Scissors = C

For you:
- X = need to lose
- Y = need to tie
- Z = need to win

What would your total score be if everything goes exactly according to your strategy guide?

*/

func main() {
	// Read the file
	input, _ := os.Open("day2input.txt")
	defer input.Close()
	sc := bufio.NewScanner(input)

	scorePart1 := 0
	scorePart2 := 0
	scoresPart1 := map[string]int{"A X": 4, "A Y": 8, "A Z": 3, "B X": 1, "B Y": 5, "B Z": 9, "C X": 7, "C Y": 2, "C Z": 6}
	scoresPart2 := map[string]int{"A X": 3, "A Y": 4, "A Z": 8, "B X": 1, "B Y": 5, "B Z": 9, "C X": 2, "C Y": 6, "C Z": 7}

	for sc.Scan() {
		scorePart1 += scoresPart1[sc.Text()]
		scorePart2 += scoresPart2[sc.Text()]
	}
	fmt.Println("Part 1:", scorePart1)
	fmt.Println("Part 2:", scorePart2)

}
