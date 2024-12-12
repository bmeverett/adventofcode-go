package main

import (
	"flag"
	"fmt"
	"github.com/bmeverett/adventofcode-go/utils"
	"strconv"
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

	stones := strings.Fields(input)

	stoneInts := make(map[int]int)
	for _, stone := range stones {
		s, err := strconv.Atoi(stone)
		if err != nil {
			panic(err)
		}
		stoneInts[s]++
	}

	cnt := 25
	if part == 2 {
		cnt = 75
	}

	for i := 0; i < cnt; i++ {
		newStones := make(map[int]int)
		for k, v := range stoneInts {
			if k == 0 {
				newStones[1] += v
			} else if len(strconv.Itoa(k))%2 == 0 {
				splt := len(strconv.Itoa(k)) / 2
				n, _ := strconv.Atoi(strconv.Itoa(k)[:splt])
				n2, _ := strconv.Atoi(strconv.Itoa(k)[splt:])

				newStones[n] += v
				newStones[n2] += v
			} else {
				newStones[k*2024] += v
			}
		}
		stoneInts = newStones
	}
	count := 0
	for _, v := range stoneInts {
		count += v
	}
	return count
}
