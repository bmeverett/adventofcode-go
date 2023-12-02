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
	
	for id, r := range lines {
		games := strings.Split(r, ":")[1]
		possible := true
		red :=0
		green :=0
		blue :=0
		for _, game := range strings.Split(games, ";") {
			cubes := strings.Split(game, ",")
			total := 0
			
			for _, cube := range cubes {
				cube = strings.Trim(cube, " ")
				num, _ := strconv.Atoi(strings.Split(cube, " ")[0])
				total += num
				color := strings.Split(cube, " ")[1]
				if color == "red" && num > 12 {
					possible = false
				} else if color == "green" && num > 13  {
					possible = false

				} else if color == "blue" && num > 14 {
					possible = false
				}

				if part == 2 {
					if color == "red" {
						red = Max(red, num)
					} else if color == "green" {
						green = Max(green, num)
					} else if color == "blue" {
						blue = Max(blue, num)
					}
				}
			}


			if total > 39 {
				possible = false
			}

			
			
		}
		

		if possible && part != 2 {
			sum += id + 1
		} else if part == 2 {
			sum += red * blue * green
		}

		
	}
	// 42817 low
	return sum
}

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

