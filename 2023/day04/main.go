package main

import (
	"flag"
	"fmt"
	"slices"
	"strings"

	"github.com/bmeverett/adventofcode-go/utils"
)

func main() {

	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()
	fmt.Println("Running part", part)

	input := utils.ReadFile("./input.txt")
	output := run(input, part)
	fmt.Println(output)

}

func run(input string, part int) int {
	sum := 0
	lines := strings.Split(input, "\n")
	totalCards := make(map[int]int)

	for init := range lines {
		totalCards[init+1] = 1
	}
	totalCount := 0
	for i, r := range lines {
		row :=  strings.Split(r, ":")
		cards := strings.Split(strings.Trim(row[1], ""), "|")

		myCard := strings.Split(strings.Trim(cards[1], " "), " ")
		winningNums := strings.Split(strings.Trim(cards[0], " "), " ")
		multipler := 0
		count := 0
		for _, num := range myCard {
			if num == "" {
				continue
			}
			if slices.Contains(winningNums, num) {
				if multipler == 0 {
					multipler = 1
					count += 1
				} else {
					multipler *=2
					count += 1
				}
			}
		}
		i += 1
		sum += multipler
		for c := i + 1; c <= count + i; c++ {
			if i == 1 {
				totalCards[c]++
				//totalCount++
			} else {
				totalCards[c] += totalCards[i]
				//totalCount+= totalCards[i]
			}
		}

		// if part == 2 && multipler == 0{
		// 	break
		// }

	}


	if part == 2 {
		for _, count := range totalCards {
			totalCount += count
		}
		return totalCount
	}
	return sum
}

