package main

import (
	"flag"
	"fmt"
	"github.com/bmeverett/adventofcode-go/utils"
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

	rows := strings.Split(input, "\n")
	count := 0
	grid := make([][]string, len(rows))

	x := 0
	y := 0
	symbol := "^"

	for i, row := range rows {
		grid[i] = make([]string, len(row))
		for j, cell := range row {
			grid[i][j] = string(cell)
			if string(cell) == "^" {
				x = i
				y = j
			}
		}
	}

	out := false
	for !out {
		if grid[x][y] == "." {
			grid[x][y] = "1"
			count++
		}
		if symbol == "^" {
			x--
		} else if symbol == ">" {
			y++
		} else if symbol == "v" {
			x++
		} else if symbol == "<" {
			y--
		}

		if x >= 0 && x < len(grid) && y >= 0 && y < len(grid) {
			if grid[x][y] == "#" {
				if symbol == "^" {
					symbol = ">"
					x++
				} else if symbol == ">" {
					symbol = "v"
					y--
				} else if symbol == "v" {
					symbol = "<"
					x--
				} else if symbol == "<" {
					symbol = "^"
					y++
				}
			}
		} else {
			out = true
		}
	}

	return count + 1
}
