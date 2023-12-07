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
	
	if part == 2{
		re := regexp.MustCompile(`\s+`)
        out := re.ReplaceAllString(lines[0], "")
        lines[0] = strings.TrimSpace(out)

		out = re.ReplaceAllString(lines[1], "")
        lines[1] = strings.TrimSpace(out)
	}

	re := regexp.MustCompile(`\d+`)
		

	times := re.FindAllString(lines[0], -1)
	distances := re.FindAllString(lines[1], -1)

	for i := 0; i<len(times); i++ {
		time, _ := strconv.Atoi(times[i])
		distance, _ := strconv.Atoi(distances[i])
		count :=0

		for t:=1; t< time; t++ {
			timeLeft := time -t 
			moved := t * timeLeft
			if moved > distance {
				count ++
			}
		} 
		if sum == 0 {
			sum = count
		} else {
			sum *= count
		}
	}

	return sum
}

