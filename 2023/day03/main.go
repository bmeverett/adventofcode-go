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

	//symbols := "!@#$%^&*()+/="
	for i :=0; i <len(grid); i ++ {
		re := regexp.MustCompile(`\d+`)
		r := grid[i]
		nums := re.FindAllString(r, -1)

		for _, num := range nums {
			index := strings.Index(r, num)
			conNum, _ := strconv.Atoi(num)
			// check left
			//strings.Contains(symbols, string(r[index-1]))
			if index - 1 > 0 && isSymbol(grid, i, index -1)   {
				sum += conNum
				break
			}

			// check right
			//strings.Contains(symbols, string(r[index + len(num)]))
			if index + len(num) < len(r) && isSymbol(grid, i, index + len(num))  {
				sum += conNum
				break
			}

			for col := index; col < col + len(num); col++ {
				if col >= len(r) {
					break
				}
				// check top
				if i > 0 && isSymbol(grid, i-1, col) {
					sum += conNum
					break
				}
				// check bottom
				if i+1 < len(r) && isSymbol(grid, i+1, col) {
					sum += conNum
					break
				}

				// check diag
				if i >0 && col > 0 && isSymbol(grid, i-1, col-1) {
					sum += conNum
					break
				}
				
				if  i +1 < len(r) && col > 0 && isSymbol(grid, i+1, col-1) {
					sum += conNum
					break
				}
				if  i >0 && col +1 < len(grid[i]) && isSymbol(grid, i-1, col +1) {
					sum += conNum
					break
				}
				if  i +1 < len(r) && col +1 < len(grid[i]) && isSymbol(grid, i+1, col+1) {
					sum += conNum
					break
				}
			}
			
		}
	}
	//352629 low
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