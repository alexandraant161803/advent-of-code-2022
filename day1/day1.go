package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	//Read the file
	input, err := os.ReadFile("day1input.txt")
	if err != nil {
		log.Fatal(err)
	}
	elfsCalories := strings.Split(string(input), "\n")

	// Part 1
	maxCalories := part1(elfsCalories)
	fmt.Println("Part 1: ", maxCalories)

	// Part 2
	topThreeSumCalories := part2(elfsCalories)
	fmt.Println("Part 2: ", topThreeSumCalories)

}

/*
Find the Elf carrying the most Calories. How many total Calories is that Elf carrying?
*/
func part1(elfsCalories []string) int {
	maxCalories := 0
	sumCalories := 0
	for _, calorie := range elfsCalories {
		if calorie == "" {
			if sumCalories > maxCalories {
				maxCalories = sumCalories
			}
			sumCalories = 0
		} else {
			calorieInt, err := strconv.Atoi(calorie)
			if err != nil {
				log.Fatalf("Could not convert %s to int", calorie)
			}
			sumCalories += calorieInt

		}

	}
	return maxCalories
}

/*
Find the top three Elves carrying the most Calories.
How many Calories are those Elves carrying in total?
*/

func part2(elfsCalories []string) int {
	sumElfsCalories := []int{}
	sum := 0
	for _, calorie := range elfsCalories {
		if calorie == "" {
			sumElfsCalories = append(sumElfsCalories, sum)
			sum = 0
		} else {
			calorieInt, err := strconv.Atoi(calorie)
			if err != nil {
				log.Fatalf("Could not convert %s to int", calorie)
			}

			sum += calorieInt
		}

	}

	// sort descending
	sort.Slice(sumElfsCalories, func(i, j int) bool {
		return sumElfsCalories[i] > sumElfsCalories[j]
	})

	return sumElfsCalories[0] + sumElfsCalories[1] + sumElfsCalories[2]

}
