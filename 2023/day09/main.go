package main

import (
	"flag"
	"fmt"
	"strconv"
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

type Instruction struct {
	Left  string
	Right string
}

func run(input string, part int) int {
	sum := 0
	lines := strings.Split(input, "\n")

	for _, r := range lines {
		vals := strings.Split(r, " ")
		pairs := make([]int, 0)
		diff := make(map[int]int)
		history := make([][]int, 0)
		for _, val := range vals {
			iVal, _ := strconv.Atoi(val)
			pairs = append(pairs, iVal)
			diff[iVal]++
		}

		history = append(history, pairs)
		for len(diff) > 1 {
			newPairs := make([]int, 0)
			clear(diff)
			diff[pairs[0]]++
			for i := 1; i < len(pairs); i++ {
				newPairs = append(newPairs, pairs[i]-pairs[i-1])
				diff[pairs[i]]++
			}

			pairs = newPairs

			history = append(history, newPairs)
		}

		for i := len(history) - 1; i >= 0; i-- {
			if part == 2 {
				if i == len(history)-1 {
					history[i] = append([]int{0}, history[i]...)
					continue
				}

				addVal := history[i][0] - history[i+1][0]
				if i == 0 {
					sum += addVal
				}
				history[i] = append([]int{addVal}, history[i]...)

			} else {
				if i == len(history)-1 {
					history[i] = append(history[i], 0)
					continue
				}

				addVal := history[i][len(history[i])-1] + history[i+1][len(history[i+1])-1]

				if i == 0 {
					sum += addVal
				}
				history[i] = append(history[i], addVal)
			}

		}

	}

	return sum

}
