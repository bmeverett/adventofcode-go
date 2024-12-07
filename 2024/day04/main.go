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
	grid := make([][]rune, len(rows))
	word := "XMAS"
	// Function to count occurrences of word in a line (both forward and backward)
	countWordInLine := func(line []rune, word string) int {
		lineStr := string(line)
		return strings.Count(lineStr, word) + strings.Count(reverseString(lineStr), word)
	}

	for i, row := range rows {
		grid[i] = make([]rune, len(row))
		for j, cell := range row {
			grid[i][j] = cell
		}

		count += countWordInLine([]rune(row), word)
	}

	for i, row := range rows {
		var vert []rune
		for j, _ := range row {
			vert = append(vert, grid[j][i])
		}

		count += countWordInLine(vert, word)
	}

	rowLen := len(rows)
	cols := len(grid[0])

	for d := -rowLen + 1; d < cols; d++ {
		var diagonal []rune
		for i := 0; i < rowLen; i++ {
			j := i - d
			if j >= 0 && j < cols {
				diagonal = append(diagonal, grid[i][j])
			}
		}
		count += countWordInLine(diagonal, word)
	}

	// Anti-diagonals (row + col constant)
	for d := 0; d < rowLen+cols-1; d++ {
		var diagonal []rune
		for i := 0; i < rowLen; i++ {
			j := d - i
			if j >= 0 && j < cols {
				diagonal = append(diagonal, grid[i][j])
			}
		}
		count += countWordInLine(diagonal, word)
	}

	return count
}

// Helper function to reverse a string
func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
