package main

import (
	"flag"
	"fmt"
	"github.com/bmeverett/adventofcode-go/utils"
	"math"
	"sort"
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

	lneSplt := strings.Split(input, "\n")
	left := make([]int, len(lneSplt))
	right := make([]int, len(lneSplt))

	lCnt := map[int]int{}
	rCnt := map[int]int{}

	for _, lne := range lneSplt {
		splt := strings.Split(lne, "   ")
		lVal, err := strconv.Atoi(splt[0])
		if err != nil {
			panic(err)
		}

		rVal, err := strconv.Atoi(splt[1])
		if err != nil {
			panic(err)
		}
		left = append(left, lVal)
		right = append(right, rVal)

		if part == 2 {
			lCnt[lVal]++
			rCnt[rVal]++
		}
	}

	sort.Ints(left)
	sort.Ints(right)

	diff := 0

	for i := 0; i < len(left); i++ {
		if part == 1 {
			diff += int(math.Abs(float64(left[i]) - float64(right[i])))
		} else {
			diff += rCnt[left[i]] * left[i]
		}
	}

	return diff
}
