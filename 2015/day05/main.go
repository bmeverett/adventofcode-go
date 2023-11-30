package main

import (
	"flag"
	"fmt"
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
	lines := strings.Split(input, "\n")
	nice := 0
	for _, r := range lines {
		if part == 1 {
			vowelCount := 0
			doubleLetter := false
			var lastLetter string
			for _, ch := range r {
				switch ch {
				case 'a', 'e', 'i', 'o', 'u':
					vowelCount++
				}
	
				if lastLetter != "" && !doubleLetter {
					doubleLetter = lastLetter == string(ch)
				} 
				lastLetter = string(ch)
			}
	
			if !strings.Contains(r, "ab") && !strings.Contains(r, "cd") && 
			   !strings.Contains(r, "pq") && !strings.Contains(r, "xy") && 
			   doubleLetter && vowelCount >= 3 {
				nice++
			}
		} else {
			count := 0
			for _, line := range lines {
				for i :=0; i< len(lines) -2; i++ {
					matchCh := lines[i: i+2]
				}
			}
		}
		
	}

	return nice
}
