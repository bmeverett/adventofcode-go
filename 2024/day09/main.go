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

	for i := len(blocks) - 1; i >= 0; i-- {
		if len(indexes) == 0 || indexes[0] >= int64(i) {
			break
		} else if blocks[i] == "." {
			continue
		}
		last := blocks[i]
		blocks[i] = "."
		blocks[indexes[0]] = last
		indexes = indexes[1:]
	}

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
	//fmt.Printf("%v\n", blocks)
	return count //4949947618371 low
}

func appendStr(fs *[]string, indexes *[]int64, s string, size int) {
	for i := 0; i < size; i++ {
		if s == "." {
			*indexes = append(*indexes, int64(len(*fs)))
		}
		*fs = append(*fs, s)

	}
}
