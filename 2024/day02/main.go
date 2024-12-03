package main

import (
	"flag"
	"fmt"
	"github.com/bmeverett/adventofcode-go/utils"
	"math"
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

	rows := strings.Split(input, "\n")
	count := 0
	bad := make([]int, 0)
	for j, row := range rows {
		line := strings.Fields(row)
		good := checkLine(line, part, &bad, j)

		if good {
			count++
		}
	}

	if part == 2 {
		for _, lne := range bad {
			str := rows[lne]
			row := strings.Fields(str)
			for i := 0; i < len(row); i++ {
				newLine := make([]string, 0)
				cpy := append(newLine, row...)
				newLine = append(cpy[:i], cpy[i+1:]...)
				dumb := make([]int, 0)
				good := checkLine(newLine, part, &dumb, 0)
				if good {
					count++
					break
				}
			}
		}

	}

	return count
}

func checkLine(line []string, part int, bad *[]int, index int) bool {
	good := false
	prev := 0
	increasing := false
	for i, lne := range line {
		iVal, err := strconv.Atoi(lne)
		if err != nil {
			panic(err)
		}
		if i == 0 {
			prev = iVal
			continue
		}

		if i == 1 {
			increasing = iVal > prev
		}

		diff := math.Abs(float64(iVal - prev))
		goodDiff := diff >= 1 && diff <= 3
		good = increasing && iVal > prev && goodDiff || !increasing && iVal < prev && goodDiff
		if !good && part == 1 {
			break
		} else if !good && part == 2 {
			*bad = append(*bad, index)
			break
		}
		prev = iVal

	}
	return good
}
