package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

/*
Find the Elf carrying the most Calories. How many total Calories is that Elf carrying?
*/

func main() {
	fmt.Println("Hello, World!")
	//Read the file
	input, _ := os.Open("day1input.txt")
	defer input.Close()
	sc := bufio.NewScanner(input)

	sc.Scan()
	maxCalories := 0
	sumCalories := 0
	for sc.Scan() {
		line := sc.Text()
		if line == "" {
			if sumCalories > maxCalories {
				maxCalories = sumCalories
			}
			sumCalories = 0
		} else {
			calorieInt, err := strconv.Atoi(line)
			if err != nil {
				panic(err)
			}
			sumCalories += calorieInt

		}

	}
	fmt.Println(maxCalories)

}
