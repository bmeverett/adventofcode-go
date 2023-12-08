package main

import (
	"flag"
	"fmt"
	"regexp"
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


type Instruction struct {
	Left string
	Right string
}

func run(input string, part int) int {
	steps := 0
	lines := strings.Split(input, "\n")
	
	nodes := make(map[string]Instruction)
	next := "AAA"
	directions := lines[0]
	directionIndex := 0
	for i:=2; i < len(lines); i++{
		ele := strings.Trim(strings.Split(lines[i], "=")[0], " ")
		re := regexp.MustCompile(`([A-Za-z]{3}),\s*([A-Za-z]{3})`)
		matches := re.FindStringSubmatch(lines[i])

		instr := Instruction{
			Left: matches[1],
			Right: matches[2],
		}

		nodes[ele] = instr

		
	}

	for next != "ZZZ" {
		steps++
		direction := string([]rune(directions)[directionIndex])
		node := nodes[next]
		if direction == "L" {
			next = node.Left
		} else if direction == "R" {
			next =node.Right
		}

		
		directionIndex++
		if directionIndex >= len(directions) {
			directionIndex = 0
		}
	}

	return steps
		
}