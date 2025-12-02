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
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()
	fmt.Println("Running part", part)

	input := utils.ReadFile("./input.txt")
	output := run(input, part)
	fmt.Println(output)
}

func run(input string, part int) int {

	lines := strings.Split(input, "\n")
	dial := 50
	zeroCount := 0
	for _, l := range lines {
		splt := strings.Split(l, "")
		dir := splt[0]
		num, _ := strconv.Atoi(strings.Join(splt[1:], ""))

		if dir == "L" {
			//fmt.Printf("L %d\n", num)
			dial -= num
			for dial < 0 {
				dial += 100
			}

		} else if dir == "R" {
			dial += num
			//fmt.Printf("R %d\n", num)
			for dial > 99 {
				dial -= 100
			}
		}

		//	fmt.Printf("Dial %d\n", dial)
		//284 low
		if dial == 0 {
			zeroCount++
		}

	}

	return zeroCount
}
