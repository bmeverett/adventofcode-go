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

func run(input string, part int) int {

	sum := 0
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		most := 0
		curr := 0
		runes := []rune(line)
		for i, l1 := range runes {
			for _, l2 := range runes[i+1:] {
				curr, _ = strconv.Atoi(string(l1) + string(l2))
				if curr > most {
					most = curr
				}
			}
		}
		fmt.Println(most)
		sum += most
	}

	return sum
}
