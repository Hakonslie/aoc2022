package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

const fileName = "chunks"

func main() {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var chunkScanner []string
	for scanner.Scan() {
		chunkScanner = append(chunkScanner, scanner.Text())
	}

	cheatsheet := make(map[string]int)
	cheatsheet["["] = 0
	cheatsheet["{"] = 1
	cheatsheet["<"] = 2
	cheatsheet["("] = 3
	cheatsheet["]"] = 4
	cheatsheet["}"] = 5
	cheatsheet[">"] = 6
	cheatsheet[")"] = 7

	additionPoints := make(map[int]int)
	additionPoints[4] = 2
	additionPoints[5] = 3
	additionPoints[6] = 4
	additionPoints[7] = 1

	var winners []int
chunkLoop:
	for _, chunk := range chunkScanner {
		var stack []int
		pointer := 0
		multiplier := 0
		for _, v := range strings.Split(chunk, "") {
			key := cheatsheet[v]
			if key < 4 {
				pointer++
				stack = append(stack, key)
			} else {
				pointer--
				if key%4 == stack[pointer] {
					stack = stack[:pointer]
				} else {
					continue chunkLoop
				}
			}
		}
		pointer = len(stack) - 1
		for pointer >= 0 {
			expectedKey := stack[pointer] + 4
			pointer--
			multiplier = multiplier * 5
			multiplier += additionPoints[expectedKey]
		}
		winners = append(winners, multiplier)
	}
	sort.Ints(winners)
	fmt.Println(winners[len(winners)/2])
}
