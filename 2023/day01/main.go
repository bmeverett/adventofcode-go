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
	for id, r := range lines {
		re := regexp.MustCompile(`\d`)
		nums := re.FindAllString(r, -1)
		var numStr string
		var first string
		var last string
	
		if len(nums) > 0 {
			first = nums[0]
			last = nums[len(nums)- 1]

		}
			
		if part == 2 {
			// commented out solution was off by one got 57346
			strNums := []string {"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
			//fmt.Printf("%s", r)
			// var foundIndexes []int
			// m := make(map[int]int)
			r = strings.ReplaceAll(r, "oneight", "18")
			r = strings.ReplaceAll(r, "twone", "21")
			r = strings.ReplaceAll(r, "threeight", "38")
			r = strings.ReplaceAll(r, "fiveight", "58")
			r = strings.ReplaceAll(r, "eighthree", "83")
			r = strings.ReplaceAll(r, "nineight", "98")
			r = strings.ReplaceAll(r, "eightwo", "82")
			r = strings.ReplaceAll(r, "sevenine", "79")
			for i, s := range strNums {
				// index := strings.Index(r, s)
				// if index != -1 {
				// 	//fmt.Printf("index %s %d\n", s, index)
				// 	// if index < strings.Index(r, first) {
				// 	// 	first = strconv.Itoa(i +1)
				// 	// }
				// 	// if index > strings.Index(r, last) {
				// 	// 	last =strconv.Itoa(i +1)
				// 	// }
				// 	foundIndexes = append(foundIndexes, index)
				// 	m[index] = i +1
				// }
				r = strings.ReplaceAll(r, s, strconv.Itoa(i +1))
			}

			nums = re.FindAllString(r, -1)
			if len(nums) > 0 {
				first = nums[0]
				last = nums[len(nums)- 1]
	
			}

			// if len(foundIndexes) > 0 {
			// 	if first != "" {
			// 		firstIndex :=strings.Index(r, first)
			// 		foundIndexes = append(foundIndexes, firstIndex)
			// 		firstVal, _ := strconv.Atoi(first)
			// 		m[firstIndex] = firstVal
	
			// 	} 
	
			// 	if last != "" {
			// 		lastIndex := strings.Index(r, last)
			// 		foundIndexes = append(foundIndexes, lastIndex)
			// 		lastVal, _ := strconv.Atoi(last)
			// 		m[lastIndex] = lastVal
			// 	}
	
			// 	sort.Ints(foundIndexes)
			// 	first = strconv.Itoa(m[foundIndexes[0]])
			// 	last = strconv.Itoa(m[foundIndexes[len(foundIndexes)-1]])
			// }
			
		}

		numStr = first + last
		num, err := strconv.Atoi(numStr)
		fmt.Println(id, first, last, num)
		if err != nil {
			fmt.Println("error parsing int")
		}
		sum += num
		
	}

	return sum
}
