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
	total := 0
	x := 0
	y := 0
	robox := 0
	roboy := 0
	grid := make(map[int]map[int]int)

	grid[x] = make(map[int]int)
	grid[x][y] = 2
	for i, r := range input {
		if part == 2 && i%2 != 0 && i > 0 {
			if r == '^' {
				robox--
			} else if r == '>' {
				roboy++
			} else if r == '<' {
				roboy--
			} else if r == 'v' {
				robox++
			}
			if _, found := grid[robox]; !found {
				grid[robox] = make(map[int]int)
			}
			grid[robox][roboy]++
		} else {
			if r == '^' {
				x--
			} else if r == '>' {
				y++
			} else if r == '<' {
				y--
			} else if r == 'v' {
				x++
			}
			if _, found := grid[x]; !found {
				grid[x] = make(map[int]int)
			}
			grid[x][y]++
		}

	}

	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] > 0 {
				total++
			}
		}

	}

	return total
}
