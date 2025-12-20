package main

import (
	"flag"
	"fmt"
	"regexp"
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

func run(input string, part int) int {
	lines := strings.Split(input, ",")
	repeatingPattern := regexp.MustCompile(`^(\d+)\1+$`)
	sum := 0
	for _, l := range lines {
		splt := strings.Split(l, "-")
		start, _ := strconv.Atoi(splt[0])
		end, _ := strconv.Atoi(splt[1])

		for i := start; i <= end; i++ {
			numStr := strconv.Itoa(i)
			if repeatingPattern.MatchString(numStr) {
				sum++
			}

		}
	}

	return sum
}
