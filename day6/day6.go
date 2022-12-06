package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	inp, err := os.ReadFile("day6input.txt")
	if err != nil {
		log.Fatal(err)
	}

	input := string(inp)
	var characters string
	for i, ch := range input {
		character := string(ch)
		characters = characters + string(character)
		if len(characters) == 4 { // change to 4 for day 1
			madeIt := isUnique(characters)
			if madeIt {
				fmt.Printf("Found it on: %v", i+1)
				break
			} else {
				characters = characters[1:]
			}
		}

	}
}

func isUnique(str string) bool {
	m := make(map[rune]bool)
	for _, i := range str {
		_, ok := m[i]
		if ok {
			return false
		}

		m[i] = true
	}

	return true
}
