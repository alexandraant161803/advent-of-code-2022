package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

/*
Info:
Lowercase item types a through z have priorities 1 through 26.
Uppercase item types A through Z have priorities 27 through 52.
*/
var priorities = map[string]int{
	"a": 1, "b": 2, "c": 3, "d": 4, "e": 5, "f": 6, "g": 7, "h": 8, "i": 9, "j": 10, "k": 11, "l": 12, "m": 13, "n": 14, "o": 15, "p": 16, "q": 17, "r": 18, "s": 19, "t": 20, "u": 21, "v": 22, "w": 23, "x": 24, "y": 25, "z": 26,
	"A": 27, "B": 28, "C": 29, "D": 30, "E": 31, "F": 32, "G": 33, "H": 34, "I": 35, "J": 36, "K": 37, "L": 38, "M": 39, "N": 40, "O": 41, "P": 42, "Q": 43, "R": 44, "S": 45, "T": 46, "U": 47, "V": 48, "W": 49, "X": 50, "Y": 51, "Z": 52}

// instead of priorities mapping one could also use ASCII value mapping: https://www.cs.cmu.edu/~pattis/15-1XX/common/handouts/ascii.html

func main() {
	input, err := os.ReadFile("day3input.txt")
	if err != nil {
		log.Fatal(err)
	}
	elfsItems := strings.Split(string(input), "\n")
	part1(elfsItems)
	part2(elfsItems)

}

/*
Task:
Find the item type that appears in both compartments of each rucksack.
What is the sum of the priorities of those item types?
*/
func part1(elfsItems []string) {
	sum := 0
	for _, compartments := range elfsItems {
		itemTypes := map[string]int{}
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
			if _, ok := itemTypes[item]; ok {
				sum += priorities[item]
				break

			}
		}

	}

	fmt.Println(sum)
}

/*

Find the one item type that is common between all three Elves in each group.
What is the sum of the priorities of those item types?
*/

func part2(elfsItems []string) {
	sum := 0
	i := 1
	itemTypes := map[string]int{}
	for _, elf := range elfsItems {
		if i == 1 {
			for _, item := range elf {
				item := string(item)
				if _, ok := itemTypes[item]; !ok {
					itemTypes[item] = 1
				}
			}

		} else if i == 2 {
			for _, item := range elf {
				item := string(item)
				if _, ok := itemTypes[item]; ok {
					itemTypes[item] = 2
				}
			}

		} else {
			i = 0
			for _, item := range elf {
				item := string(item)
				if _, ok := itemTypes[item]; ok {
					if itemTypes[item] == 2 {
						sum += priorities[item]
						break
					}
				}
			}
			itemTypes = map[string]int{}
		}
		i++

	}

	fmt.Println(sum)
}
