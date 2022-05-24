package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const fileName = "vents"
const fileNameExample = "ventsexample"

func main() {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	type position struct {
		x int
		y int
	}
	type line struct {
		from position
		to   position
	}
	var lines []line

	for scanner.Scan() {
		numbers := scanner.Text()
		vectorSplit := strings.Split(numbers, " -> ")

		var newLine line

		for i, pos := range vectorSplit {

			posSplit := strings.Split(pos, ",")
			xPos, _ := strconv.Atoi(posSplit[0])
			yPos, _ := strconv.Atoi(posSplit[1])
			if i == 0 {
				newLine.from = position{xPos, yPos}
			} else {
				newLine.to = position{xPos, yPos}
			}
		}
		lines = append(lines, newLine)
	}

	var board [990][990]int

	for _, l := range lines {
		fromX := l.from.x
		fromY := l.from.y
		toX := l.to.x
		toY := l.to.y

		board[fromY][fromX]++
		// While from and to are not the same
		for fromX != toX || fromY != toY {

			// Diagonal difference
			if fromY != toY && fromX != toX {
				if fromY < toY {
					fromY++
				} else {
					fromY--
				}
				if fromX < toX {
					fromX++
				} else {
					fromX--
				}
				// Vertical difference
			} else if fromY != toY {
				if fromY > toY {
					fromY--
				} else if fromY < toY {
					fromY++
				}
				// Horizontal difference
			} else if fromX != toX {
				if fromX > toX {
					fromX--
				} else if fromX < toX {
					fromX++
				}
			}
			board[fromY][fromX]++
		}
	}

	overlaps := 0

	for _, o := range board {
		fmt.Println(o)
		for _, p := range o {
			if p > 1 {
				overlaps++
			}
		}
	}
	fmt.Println(overlaps)

}
