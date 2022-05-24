package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func indexOf(element string, data []string) int {
	for k, v := range data {
		if element == v {
			return k
		}
	}
	return -1 //not found.
}

func main() {
	file, err := os.Open("bingoexample")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	_ = scanner.Scan()
	numbers := scanner.Text()
	numbersSplit := strings.Split(numbers, ",")

	inputs := make(map[int][]string)

	i := 0
	for scanner.Scan() {
		newInput := scanner.Text()
		if newInput == "" {
			i++
		} else {
			inputArray := strings.Split(newInput, " ")
			for _, inp := range inputArray {
				if inp != "" {
					inputs[i] = append(inputs[i], inp)
				}
			}
		}
	}
	winningConditions := [10][5]int{
		{0, 1, 2, 3, 4},
		{5, 6, 7, 8, 9},
		{10, 11, 12, 13, 14},
		{15, 16, 17, 18, 19},
		{20, 21, 22, 23, 24},
		{0, 5, 10, 15, 20},
		{1, 6, 11, 16, 21},
		{2, 7, 12, 17, 22},
		{3, 8, 13, 18, 23},
		{4, 9, 14, 19, 24},
	}
	topBoardIterationsReq := 0
	var topBoardIndexes map[int]bool
	var topBoardIndex int
	var topBoardLastNumber int

	//mainLoop iterates one board after the other
mainLoop:
	for boardIndex, elem := range inputs {
		// indexes of marked numbers
		indexes := make(map[int]bool)

		// Iterate through randomly picked numbers
		for numberIndex, n := range numbersSplit {
			// get index of one randomly picked number
			ind := indexOf(n, elem)
			// is it in the current board(elem)?
			if ind != -1 {
				indexes[ind] = true
			}

			// Iterate over winning conditions
		conditionLoop:
			for _, winningElem := range winningConditions {
				for _, h := range winningElem {
					// Per winning condition, if there is no match. go to next winning condition
					if !indexes[h] {
						continue conditionLoop
					}
				}
				if numberIndex > topBoardIterationsReq {
					topBoardIterationsReq = numberIndex
					topBoardIndexes = indexes
					topBoardIndex = boardIndex
					topBoardLastNumber, _ = strconv.Atoi(n)
				}
				// if it got points, continue mainloop
				continue mainLoop

			}
		}
	}

	sumOfUnmarked := 0
	for i, o := range inputs[topBoardIndex] {
		if !topBoardIndexes[i] {
			atoi, _ := strconv.Atoi(o)
			sumOfUnmarked += atoi
		}
	}
	fmt.Println(sumOfUnmarked * topBoardLastNumber)
}
