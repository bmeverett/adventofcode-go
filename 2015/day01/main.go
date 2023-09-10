package main

import (
	"flag"
	"fmt"

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
	floor := 0
	for i, r := range input {
		if r == '(' {
			floor += 1
		} else if r == ')' {
			floor -= 1
		}

		if part == 2 && floor == -1 {
			return i + 1
		}
	}

	return floor
}
