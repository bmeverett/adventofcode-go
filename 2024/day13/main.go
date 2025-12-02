package main

import (
	"flag"
	"fmt"
	"github.com/bmeverett/adventofcode-go/utils"
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

	count := 0
	for _, line := range strings.Split(input, "\n") {
		var ax, ay, bx, by, px, py int
		if strings.Contains(line, "Button A") {

		} else if strings.Contains(line, "Button B") {

		} else if strings.Contains(line, "Prize") {

		} else {
			a = ""
			b = ""
			continue
		}
	}
	return count
}
