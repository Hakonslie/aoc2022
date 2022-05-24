package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

	illegalPoints := make(map[int]int)
	illegalPoints[4] = 57
	illegalPoints[5] = 1197
	illegalPoints[6] = 25137
	illegalPoints[7] = 3

	syntaxError := 0
chunkLoop:
	for _, chunk := range chunkScanner {
		var stack []int
		pointer := 0
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
					syntaxError += illegalPoints[key]
					continue chunkLoop
				}
			}
		}
	}
	fmt.Println(syntaxError)
}
