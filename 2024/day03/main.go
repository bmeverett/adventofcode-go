package main

import (
	"flag"
	"fmt"
	"github.com/bmeverett/adventofcode-go/utils"
	"regexp"
	"strconv"
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

	if part == 1 {
		return part1(input)
	}

	return part2(input)
}

func part1(input string) int {
	count := 0
	re := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)
	numRe := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)

	// Find all matches
	matches := re.FindAllString(input, -1)

	// Print the matches
	for _, match := range matches {
		submatches := numRe.FindStringSubmatch(match)
		n1, err := strconv.Atoi(submatches[1])
		if err != nil {
			panic(err)
		}
		n2, err := strconv.Atoi(submatches[2])
		if err != nil {
			panic(err)
		}
		count += n1 * n2
	}
	return count
}

func part2(input string) int {
	count := 0
	// Regex to match valid mul(X,Y) instructions
	mulRe := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)

	// Regex to match do() and don't() instructions
	controlRe := regexp.MustCompile(`(do\(\)|don't\(\))`)

	// Initial state: mul instructions are enabled
	isEnabled := true

	// Find all occurrences of mul and control instructions in order
	allMatches := controlRe.FindAllStringIndex(input, -1)
	allMulMatches := mulRe.FindAllStringSubmatchIndex(input, -1)

	// Combine control and mul matches in order of appearance
	matches := append(allMatches, allMulMatches...)
	matches = sortByStart(matches)
	for _, match := range matches {
		if match[0] >= len(input) {
			continue
		}
		text := input[match[0]:match[1]]

		// Check for do() and don't() instructions
		if text == "do()" {
			isEnabled = true
		} else if text == "don't()" {
			isEnabled = false
		} else if mulMatch := mulRe.FindStringSubmatch(text); len(mulMatch) == 3 && isEnabled {
			// Perform multiplication if enabled
			x, _ := strconv.Atoi(mulMatch[1])
			y, _ := strconv.Atoi(mulMatch[2])
			count += x * y
		}
	}

	return count
}

func sortByStart(matches [][]int) [][]int {
	for i := 0; i < len(matches)-1; i++ {
		for j := 0; j < len(matches)-i-1; j++ {
			if matches[j][0] > matches[j+1][0] {
				matches[j], matches[j+1] = matches[j+1], matches[j]
			}
		}
	}
	return matches
}
