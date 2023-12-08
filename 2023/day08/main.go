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
	Left  string
	Right string
}

func run(input string, part int) int {
	steps := 0
	lines := strings.Split(input, "\n")

	nodes := make(map[string]Instruction)
	next := "AAA"
	directions := lines[0]
	directionIndex := 0
	aLines := make([]string, 0)
	for i := 2; i < len(lines); i++ {
		ele := strings.Trim(strings.Split(lines[i], "=")[0], " ")
		re := regexp.MustCompile(`([A-Za-z0-9]{3}),\s*([A-Za-z0-9]{3})`)
		matches := re.FindStringSubmatch(lines[i])

		instr := Instruction{
			Left:  matches[1],
			Right: matches[2],
		}

		nodes[ele] = instr
		if strings.HasSuffix(ele, "A") {
			aLines = append(aLines, ele)
		}

	}

	if part == 2 {
		//zCount := 0
		//direction := string([]rune(directions)[directionIndex])
		stepsMultiplier := []int{}
		//for zCount < len(aLines) {
			//newNodes := make([]string, 0)
			for _, a := range aLines {
				steps = 0
				next = a
				for !strings.HasSuffix(next, "Z") {
					steps++
					direction := string([]rune(directions)[directionIndex])
					node := nodes[next]
					if direction == "L" {
						next = node.Left
					} else if direction == "R" {
						next = node.Right
					}
			
					directionIndex++
					if directionIndex >= len(directions) {
						directionIndex = 0
					}
				}

				stepsMultiplier = append(stepsMultiplier, steps)
				// node := nodes[a]
				// if direction == "L" {
				// 	newNodes = append(newNodes, node.Left)
				// 	if strings.HasSuffix(node.Left, "Z") {
				// 		zCount++
				// 	}
				// } else if direction == "R" {
				// 	newNodes = append(newNodes, node.Right)
				// 	if strings.HasSuffix(node.Right, "Z") {
				// 		zCount++
				// 	}
				// }
			}

			// directionIndex++
			// if directionIndex >= len(directions) {
			// 	directionIndex = 0
			// }
			// aLines = newNodes
		//}
		fmt.Println(stepsMultiplier)
		result := stepsMultiplier[0]
		for i := 1; i < len(stepsMultiplier); i++ {
            result = lcm(result, stepsMultiplier[i])
        }
		return result
	}
	for next != "ZZZ" {
		steps++
		direction := string([]rune(directions)[directionIndex])
		node := nodes[next]
		if direction == "L" {
			next = node.Left
		} else if direction == "R" {
			next = node.Right
		}

		directionIndex++
		if directionIndex >= len(directions) {
			directionIndex = 0
		}
	}

	return steps

}

func gcd(a, b int) int {
    for b != 0 {
        t := b
        b = a % b
        a = t
    }

    return a
}

func lcm(a, b int) int {
    return a * b / gcd(a, b)
}

