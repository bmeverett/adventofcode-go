package main

import (
	"flag"
	"fmt"
	"github.com/bmeverett/adventofcode-go/utils"
	"strconv"
	"strings"
)

func main() {

	var part int
	flag.IntVar(&part, "part", 2, "part 1 or 2")
	flag.Parse()
	fmt.Println("Running part", part)

	input := utils.ReadFile("./input.txt")
	output := run(input, part)
	fmt.Println(output)

}

func run(input string, part int) int {

	rows := strings.Split(input, "\n")
	count := 0
	for _, row := range rows {
		splt := strings.Split(row, ":")
		sum, err := strconv.Atoi(splt[0])

		if err != nil {
			panic(err)
		}

		nums := strings.Fields(splt[1])
		ints := make([]int, 0)
		for _, num := range nums {
			n, err := strconv.Atoi(num)
			if err != nil {
				panic(err)
			}
			ints = append(ints, n)
		}
		n := len(nums) - 1

		result := GenerateCombinations(n, part)

		for _, res := range result {
			trySum := ints[0]
			for i := 0; i < len(res); i++ {
				if res[i] == "p" {
					trySum += ints[i+1]
				} else if res[i] == "x" {
					trySum *= ints[i+1]
				} else if res[i] == "|" {
					comb := strconv.Itoa(trySum) + strconv.Itoa(ints[i+1])
					trySum, err = strconv.Atoi(comb)
					if err != nil {
						panic(err)
					}
				}
			}
			if trySum == sum {
				count += sum
				break
			}
		}
	}

	return count
}

func GenerateCombinations(n, part int) [][]string {
	var result [][]string
	var combination []string

	// Helper function to generate combinations recursively
	var backtrack func(index int)
	backtrack = func(index int) {
		if index == n {
			// Append a copy of the current combination to the result
			combinationCopy := make([]string, n)
			copy(combinationCopy, combination)
			result = append(result, combinationCopy)
			return
		}
		// Choose p
		combination = append(combination, "p")
		backtrack(index + 1)
		combination = combination[:len(combination)-1]

		// Choose x
		combination = append(combination, "x")
		backtrack(index + 1)
		combination = combination[:len(combination)-1]

		if part == 2 {
			combination = append(combination, "|")
			backtrack(index + 1)
			combination = combination[:len(combination)-1]
		}
	}

	backtrack(0)
	return result
}
