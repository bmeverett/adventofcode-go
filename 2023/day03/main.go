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
	sum := 0
	lines := strings.Split(input, "\n")

	grid := make(map[int]string)
	for i, r := range lines {
		grid[i] = r
	}

	for i := 0; i < len(grid); i++ {
		re := regexp.MustCompile(`\d+`)
		r := grid[i]
		nums := re.FindAllString(r, -1)
		multiples := 0
		for _, num := range nums {
			found := false
			index := strings.Index(r, num)
			iReg := regexp.MustCompile(`\D` + num + `\D`)
			nIndex := iReg.FindAllStringIndex(r, -1)
			if nIndex != nil {
				if len(nIndex) == 1 {
					index = nIndex[0][0] + 1
				} else {
					index = nIndex[multiples][0] + 1
					multiples++
				}
			}
			singleIndex := 0
			if len(num) == 1 {
				re := regexp.MustCompile(`\b\d\b`)
				singles := re.FindStringIndex(r)
				index = singles[singleIndex]
				singleIndex++
			}

			conNum, _ := strconv.Atoi(num)
			// check left
			if index-1 > 0 && isSymbol(grid, i, index-1) {
				sum += conNum
				continue
			}

			// check right
			if index+len(num) < len(r) && isSymbol(grid, i, index+len(num)) {
				sum += conNum
				continue
			}

			for col := index; col < index+len(num); col++ {
				if col >= len(r) {
					break
				}
				// check top
				if i > 0 && isSymbol(grid, i-1, col) {
					found = true
				}

				// check bottom
				if i+1 < len(r) && isSymbol(grid, i+1, col) {
					found = true
				}

				// check diag
				if i > 0 && col > 0 && isSymbol(grid, i-1, col-1) || //top left
					i+1 < len(r) && col > 0 && isSymbol(grid, i+1, col-1) || // bottom left
					i > 0 && col+1 < len(grid[i]) && isSymbol(grid, i-1, col+1) || // top right
					i+1 < len(r) && col+1 < len(grid[i]) && isSymbol(grid, i+1, col+1) /*bottom right*/ { 
					found = true
				}

			}
			if found {
				sum += conNum
			}
		}

	}

	return sum
}

func isSymbol(grid map[int]string, i int, col int) bool {
	cell := string(grid[i][col])
	_, err := strconv.Atoi(cell)
	if cell != "." && err != nil {
		return true
	}
	return false
}
