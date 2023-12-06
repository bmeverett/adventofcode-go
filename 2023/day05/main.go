package main

import (
	"flag"
	"fmt"
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
	//sum := 0
	lines := strings.Split(input, "\n")

	var seeds []int
	seedSoil := make(map[int]int)
	soilFertilizer := make(map[int]int)
	fertilizerWater := make(map[int]int)
	waterLight := make(map[int]int)
	lightTmp := make(map[int]int)
	tmpHum := make(map[int]int)
	humLoc := make(map[int]int)

	i := 0
	seedStr := strings.Trim(strings.Split(lines[i], ":")[1], " ")
	strSplit := strings.Split(seedStr, " ")
	if part == 1 {

		for _, seed := range strSplit {
			seedInt, _ := strconv.Atoi(seed)
			seeds = append(seeds, seedInt)
		}
	} else {
		return part2(strSplit, seedStr, seeds, lines)
	}
	minLoc := 0
	for _, seed := range seeds {
		var soil, fertilizer, water, light, tmp, hum, loc int
		var ok bool
		i = 2
		for i < len(lines) {
			if lines[i] == "" {
				i++
				continue
			}

			if strings.Contains(lines[i], ":") {
				if strings.Contains(lines[i], "seed-to-soil") {
					populateMap(seedSoil, lines, &i, seed)
					soil, ok = seedSoil[seed]
					if !ok {
						soil = seed
					}
				} else if strings.Contains(lines[i], "soil-to-fertilizer") {
					populateMap(soilFertilizer, lines, &i, soil)
					fertilizer, ok = soilFertilizer[soil]
					if !ok {
						fertilizer = soil
					}
				} else if strings.Contains(lines[i], "fertilizer-to-water") {
					populateMap(fertilizerWater, lines, &i, fertilizer)
					water, ok = fertilizerWater[fertilizer]
					if !ok {
						water = fertilizer
					}
				} else if strings.Contains(lines[i], "water-to-light") {
					populateMap(waterLight, lines, &i, water)
					light, ok = waterLight[water]
					if !ok {
						light = water
					}
				} else if strings.Contains(lines[i], "light-to-temperature") {
					populateMap(lightTmp, lines, &i, light)
					tmp, ok = lightTmp[light]
					if !ok {
						tmp = light
					}
				} else if strings.Contains(lines[i], "temperature-to-humidity") {
					populateMap(tmpHum, lines, &i, tmp)
					hum, ok = tmpHum[tmp]
					if !ok {
						hum = tmp
					}
				} else if strings.Contains(lines[i], "humidity-to-location") {
					populateMap(humLoc, lines, &i, hum)
					loc, ok = humLoc[hum]
					if !ok {
						loc = hum
					}
				}
			} else {
				i++
			}
		}
		if minLoc == 0 {
			minLoc = loc
		} else if loc < minLoc {
			minLoc = loc
		}

	}

	return minLoc
}

func part2(strSplit []string, seedStr string, seeds []int, lines []string) int {
	idx := 0
	seedSoil := make(map[int]int)
	soilFertilizer := make(map[int]int)
	fertilizerWater := make(map[int]int)
	waterLight := make(map[int]int)
	lightTmp := make(map[int]int)
	tmpHum := make(map[int]int)
	humLoc := make(map[int]int)
	minLoc := 0
	 
	for idx < len(strSplit) {
		seedInt, _ := strconv.Atoi(strSplit[idx])
		if idx+1 < len(seedStr) {
			cnt, _ := strconv.Atoi(strSplit[idx+1])
			for rngeIdx := 0; rngeIdx < cnt-1; rngeIdx++ {
				seeds = append(seeds, seedInt+rngeIdx)
			}
		}
		idx += 2
		fmt.Println(idx)
		for _, seed := range seeds {
			var soil, fertilizer, water, light, tmp, hum, loc int
			var ok bool
			i := 2
			for i < len(lines) {
				if lines[i] == "" {
					i++
					continue
				}

				if strings.Contains(lines[i], ":") {
					if strings.Contains(lines[i], "seed-to-soil") {
						populateMap(seedSoil, lines, &i, seed)
						soil, ok = seedSoil[seed]
						if !ok {
							soil = seed
						}
					} else if strings.Contains(lines[i], "soil-to-fertilizer") {
						populateMap(soilFertilizer, lines, &i, soil)
						fertilizer, ok = soilFertilizer[soil]
						if !ok {
							fertilizer = soil
						}
					} else if strings.Contains(lines[i], "fertilizer-to-water") {
						populateMap(fertilizerWater, lines, &i, fertilizer)
						water, ok = fertilizerWater[fertilizer]
						if !ok {
							water = fertilizer
						}
					} else if strings.Contains(lines[i], "water-to-light") {
						populateMap(waterLight, lines, &i, water)
						light, ok = waterLight[water]
						if !ok {
							light = water
						}
					} else if strings.Contains(lines[i], "light-to-temperature") {
						populateMap(lightTmp, lines, &i, light)
						tmp, ok = lightTmp[light]
						if !ok {
							tmp = light
						}
					} else if strings.Contains(lines[i], "temperature-to-humidity") {
						populateMap(tmpHum, lines, &i, tmp)
						hum, ok = tmpHum[tmp]
						if !ok {
							hum = tmp
						}
					} else if strings.Contains(lines[i], "humidity-to-location") {
						populateMap(humLoc, lines, &i, hum)
						loc, ok = humLoc[hum]
						if !ok {
							loc = hum
						}
					}
				} else {
					i++
				}
			}
			if minLoc == 0 {
				minLoc = loc
			} else if loc < minLoc {
				minLoc = loc
			}
		}
	}
	return minLoc
}

func populateMap(theMap map[int]int, lines []string, i *int, seed int) {
	for lines[*i] != "" && *i < len(lines) {
		*i++
		if *i >= len(lines) {
			break
		}
		var dest, source, rnge int
		line := strings.Split(lines[*i], " ")
		if len(line) < 3 {
			continue
		}
		dest, _ = strconv.Atoi(line[0])
		source, _ = strconv.Atoi(line[1])
		rnge, _ = strconv.Atoi(line[2])

		if source <= seed && seed <= source+rnge {
			diff := seed - source
			theMap[seed] = dest + diff
			break
		}
	}

}
