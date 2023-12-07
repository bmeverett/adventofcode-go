package main

import (
	"flag"
	"fmt"
	"sort"
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

var cardVal = map[string]int{
	"A": 14,
	"K": 13,
	"Q": 12,
	"J": 11,
	"T": 10,
	"9": 9,
	"8": 8,
	"7": 7,
	"6": 6,
	"5": 5,
	"4": 4,
	"3": 3,
	"2": 2,
}

func run(input string, part int) int {
	sum := 0
	lines := strings.Split(input, "\n")
	five := make([]string, 0)
	four := make([]string, 0)
	full := make([]string, 0)
	three := make([]string, 0)
	two := make([]string, 0)
	one := make([]string, 0)
	high := make([]string, 0)

	handToVal := make(map[string]int)

	for _, r := range lines {
		line := strings.Split(r, " ")
		hand := line[0]
		val, _ := strconv.Atoi(line[1])
		handToVal[hand] = val
		cardCount := make(map[string]int)
		values := make([]int, 0)
		for i := 0; i < len(hand); i++ {
			card := string([]rune(hand)[i])
			_, exists := cardCount[card]
			if exists {
				cardCount[card]++
			} else {
				cardCount[card] = 1
			}
		}

		for c,v := range cardCount {
			if c == "J" && part ==2 {
				continue
			}
			values = append(values, v)
		}

		sort.Ints(values)

		length := len(cardCount)
		
		_, hasJack := cardCount["J"]
		if part == 2 && hasJack {
			cardVal["J"] = 0
			jacks := cardCount["J"]
			delete(cardCount, "J")
			length = len(cardCount)
			if length == 1 || length == 0 {
				// 5 of a kind
				five = append(five, hand)
			} else if length == 2 {
				// 4 of a kind or full house
				twoCount := 0
				for _, cnt := range values {
					if cnt == 3 || cnt == 1 {
						four = append(four, hand)
						break
					} else if cnt == 2 {
						twoCount++
					}
				}

				if twoCount == 2 {
					full = append(full, hand)
				} else if twoCount == 1 && jacks == 2 {
					four = append(three, hand)
				} else if twoCount == 1 && jacks == 1 {
					three = append(three, hand)
				}
			} else if length == 3 {
				// 3 of a kind or two pair
				threeTwoCount := 0
				for _, cnt := range values {
					if cnt == 2 {
						threeTwoCount++
					} 
				}

				if threeTwoCount == 2 {
					full = append(full, hand)
				} else {
					three = append(three, hand)
				}
			} else if length == 4 {
				// 1 pair
				one = append(one, hand)
			} 
		} else {
			if length == 1 {
				// 5 of a kind
				five = append(five, hand)
			} else if length == 2 {
				// 4 of a kind or full house
				for _, cnt := range values {
					if cnt == 4 || cnt == 1 {
						four = append(four, hand)
						break
					} else if cnt == 3 || cnt == 2 {
						full = append(full, hand)
						break
					}
				}
			} else if length == 3 {
				// 3 of a kind or two pair
				for _, cnt := range values {
					if cnt == 3 {
						three = append(three, hand)
						break
					} else if cnt == 2 {
						two = append(two, hand)
						break
					}
				}
			} else if length == 4 {
				// 1 pair
				one = append(one, hand)
			} else {
				// high card
				high = append(high, hand)
			}
		}
	}

	sortHands(high)
	sortHands(one)
	sortHands(two)
	sortHands(full)
	sortHands(three)
	sortHands(four)
	sortHands(five)

	finalIndex := 1
	for _, f := range high {
		sum += handToVal[f] * finalIndex
		finalIndex++
	}

	for _, f := range one {
		sum += handToVal[f] * finalIndex
		finalIndex++
	}

	for _, f := range two {
		sum += handToVal[f] * finalIndex
		finalIndex++
	}

	for _, f := range three {
		sum += handToVal[f] * finalIndex
		finalIndex++
	}

	for _, f := range full {
		sum += handToVal[f] * finalIndex
		finalIndex++
	}

	for _, f := range four {
		sum += handToVal[f] * finalIndex
		finalIndex++
	}

	for _, f := range five {
		sum += handToVal[f] * finalIndex
		finalIndex++
	}

	// 253521250 low
	// 288675511 high
	return sum
}

func sortHands(arr []string) {
	sort.SliceStable(arr, func(i, j int) bool {
		for ch := 0; ch < len(arr[i]); ch++ {
			left := string([]rune(arr[i])[ch])
			right := string([]rune(arr[j])[ch])

			if cardVal[left] == cardVal[right] {
				continue
			}
			return cardVal[left] < cardVal[right]
		}
		return false
	})
}
