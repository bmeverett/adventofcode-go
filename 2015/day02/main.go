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
	total := 0
	lines := strings.Split(input, "\n")
	for _, r := range lines {
		demnsions := strings.Split(r, "x")
		l, _ := strconv.Atoi(demnsions[0])
		w, _ := strconv.Atoi(demnsions[1])
		h, _ := strconv.Atoi(demnsions[2])

		if part == 1 {
			total += 2*l*w + 2*w*h + 2*h*l

			side1 := l * w
			side2 := w * h
			side3 := l * h

			minSide := min(min(side1, side2), side3)
			total += minSide
		} else if part == 2 {
			p1 := 2*l + 2*w
			p2 := 2*w + 2*h
			p3 := 2*l + 2*h

			minPar := min(min(p1, p2), p3)
			volume := l * w * h
			total += minPar + volume
		}

	}

	return total
}
