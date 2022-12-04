package main

import (
	"bufio"
	"fmt"
	"os"
)

/*In how many assignment pairs does one range fully contain the other?*/
func main() {
	input, _ := os.Open("day4input.txt")
	defer input.Close()
	sc := bufio.NewScanner(input)

	var sum int
	for sc.Scan() {
		var minPair1, maxPair1, minPair2, maxPair2 int
		fmt.Sscanf(sc.Text(), "%v-%v,%v-%v", &minPair1, &maxPair1, &minPair2, &maxPair2)

		if minPair2 > maxPair1 || maxPair2 < minPair1 {
			continue
		}

		sum++

	}
	fmt.Println(sum)
}
