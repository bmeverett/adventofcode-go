package main

import (
	"flag"
	"fmt"
	"github.com/bmeverett/adventofcode-go/utils"
	"strconv"
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

func run(input string, part int) int64 {

	count := int64(0)
	blocks := make([]string, 0)
	indexes := make([]int64, 0)
	file := true
	id := 0
	for _, c := range input {
		str := string(c)
		num, err := strconv.Atoi(str)
		if err != nil {
			panic(err)
		}
		if file {
			appendStr(&blocks, &indexes, strconv.Itoa(id), num)
			id++
		} else {
			appendStr(&blocks, &indexes, ".", num)
		}
		file = !file
	}

	fileCount := 0
	currentFile := ""
	fmt.Printf("%v\n", blocks)

	spaces := 1
	for i := len(blocks) - 1; i >= 0; i-- {
		if len(indexes) == 0 || indexes[0] >= int64(i) {
			break
		} else if blocks[i] == "." {
			spaces++
			continue
		}
		if part == 1 {
			last := blocks[i]
			blocks[i] = "."
			blocks[indexes[0]] = last
			indexes = indexes[1:]
		} else {
			if currentFile == "" {
				currentFile = blocks[i]
				fileCount++
			} else if currentFile == blocks[i] {
				fileCount++
			} else {
				//prev := int64(0)
				//iCount := 0
				//iMap := make(map[int]int64)
				//for j, index := range indexes {
				//	if j == 0 {
				//		prev = index
				//		iMap[j] = index
				//		iCount++
				//		continue
				//	}
				//
				//	if index-prev > 1 {
				//		if fileCount <= iCount {
				//			break
				//		} else {
				//			iMap = make(map[int]int64)
				//			iCount = 0
				//		}
				//	}
				//	iMap[j] = index
				//	iCount++
				//	prev = index
				//
				//}
				bCount := 0
				open := make([]int, 0)
				for j, b := range blocks {
					if b == "." {
						bCount++
						open = append(open, j)
					} else if bCount > 0 && bCount <= fileCount {
						// move
						for k, o := range open {
							blocks[o] = currentFile
							blocks[1+i+k] = "."
						}
						bCount = 0
					} else {
						bCount = 0
					}

				}

				// move file
				//if fileCount <= iCount {
				//	//for j := 0; j < fileCount; j++ {
				//	//	blocks[iMap[j]] = currentFile
				//	//	blocks[i+1+j] = "."
				//	//	//indexes = indexes[1:]
				//	//}
				//
				//	cnt := 0
				//	keys := make([]int, 0)
				//	for k, _ := range iMap {
				//		keys = append(keys, k)
				//	}
				//	sort.Ints(keys)
				//	for _, k := range keys {
				//		if cnt+1+i >= len(blocks) {
				//			break
				//		}
				//		blocks[iMap[k]] = currentFile
				//		blocks[cnt+spaces+i] = "."
				//		indexes = append(indexes[:k], indexes[k+1:]...)
				//		cnt++
				//	}
				//}
				fmt.Printf("%v\n", blocks)
				currentFile = blocks[i]
				fileCount = 1
				//iMap = make(map[int]int64)
				spaces = 1

			}
		}
	}

	fmt.Printf("%v\n", blocks)

	for i, block := range blocks {
		if block == "." {
			break
		}
		b, err := strconv.ParseInt(block, 10, 32)
		if err != nil {
			panic(err)
		}
		count += b * int64(i)
	}

	return count
}

func part2(blocks []string) {
	for i := len(blocks) - 1; i >= 0; i-- {
		for j := 0; j < len(blocks); j++ {

		}
	}
}

func appendStr(fs *[]string, indexes *[]int64, s string, size int) {
	for i := 0; i < size; i++ {
		if s == "." {
			*indexes = append(*indexes, int64(len(*fs)))
		}
		*fs = append(*fs, s)

	}
}
