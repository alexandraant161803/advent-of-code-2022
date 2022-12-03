package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
Info:
Lowercase item types a through z have priorities 1 through 26.
Uppercase item types A through Z have priorities 27 through 52.

Task:
Find the item type that appears in both compartments of each rucksack.
What is the sum of the priorities of those item types?

*/

func main() {
	// Read the file
	input, _ := os.Open("day3input.txt")
	defer input.Close()
	sc := bufio.NewScanner(input)

	priorities := map[string]int{
		"a": 1, "b": 2, "c": 3, "d": 4, "e": 5, "f": 6, "g": 7, "h": 8, "i": 9, "j": 10, "k": 11, "l": 12, "m": 13, "n": 14, "o": 15, "p": 16, "q": 17, "r": 18, "s": 19, "t": 20, "u": 21, "v": 22, "w": 23, "x": 24, "y": 25, "z": 26,
		"A": 27, "B": 28, "C": 29, "D": 30, "E": 31, "F": 32, "G": 33, "H": 34, "I": 35, "J": 36, "K": 37, "L": 38, "M": 39, "N": 40, "O": 41, "P": 42, "Q": 43, "R": 44, "S": 45, "T": 46, "U": 47, "V": 48, "W": 49, "X": 50, "Y": 51, "Z": 52}

	sum := 0
	for sc.Scan() {
		itemTypes := map[string]int{}
		compartments := sc.Text()
		firstCompartment := compartments[:len(compartments)/2]
		secondCompartment := compartments[len(compartments)/2:]

		for _, item := range firstCompartment {
			item := string(item)
			if _, ok := itemTypes[item]; !ok {
				itemTypes[item] = 0
			}
		}

		for _, item := range secondCompartment {
			item := string(item)
			if occurences, ok := itemTypes[item]; ok {
				if occurences >= 1 {
					continue
				}
				sum += priorities[item]
				itemTypes[item] += 1

			}
		}

	}

	fmt.Println(sum)
}
