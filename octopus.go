package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const fileName = "octopusExample"

type coord struct {
	row, col int
}

type Octopus struct {
	didFlash bool
	energy   int
	adjacent []coord
}

func flash(row, col int) (flashes int) {
	flashes = 0
	if !oMap[row][col].didFlash && oMap[row][col].energy > 9 {
		oMap[row][col].didFlash = true
		oMap[row][col].energy = 0
		for _, v := range oMap[row][col].adjacent {
			flashes += energize(v.row, v.col, false)
		}
		flashes++
	}
	return
}
func energize(row, col int, refuel bool) (flashes int) {
	flashes = 0
	if !oMap[row][col].didFlash {
		oMap[row][col].energy++
		if oMap[row][col].energy > 9 && !refuel {
			flashes += flash(row, col)
		}
	}
	return
}

var oMap [10][10]Octopus

func main() {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	row := 0
	for scanner.Scan() {
		col := 0
		for _, v := range strings.Split(scanner.Text(), "") {
			o, _ := strconv.Atoi(v)
			var oc Octopus
			oc.didFlash = false
			oc.energy = o
			var c []coord
			oc.adjacent = c
			if row != 0 && col != 0 {
				oc.adjacent = append(oc.adjacent, coord{row: row - 1, col: col - 1})
			}
			if row != 0 && col != 9 {
				oc.adjacent = append(oc.adjacent, coord{row: row - 1, col: col + 1})
			}
			if row != 9 && col != 0 {
				oc.adjacent = append(oc.adjacent, coord{row: row + 1, col: col - 1})
			}
			if row != 9 && col != 9 {
				oc.adjacent = append(oc.adjacent, coord{row: row + 1, col: col + 1})
			}
			if row != 0 {
				oc.adjacent = append(oc.adjacent, coord{row: row - 1, col: col})
			}
			if row != 9 {
				oc.adjacent = append(oc.adjacent, coord{row: row + 1, col: col})
			}
			if col != 0 {
				oc.adjacent = append(oc.adjacent, coord{row: row, col: col - 1})
			}
			if col != 9 {
				oc.adjacent = append(oc.adjacent, coord{row: row, col: col + 1})
			}
			oMap[row][col] = oc
			col++
		}
		row++
	}

	flashes := 0
	gameSteps := 0
	for gameSteps < 400 {
		// Reset all Octopi and increase energy levels by 1
		for i, v := range oMap {
			for j, _ := range v {
				oMap[i][j].didFlash = false
				energize(i, j, true)
			}
		}

		// Keep flashing until noone flashed
		flashesThisStep := 0
		someoneFlashed := true
		for someoneFlashed {
			someoneFlashed = false
			for i, v := range oMap {
				for j, _ := range v {
					amountOfFlashes := flash(i, j)
					if amountOfFlashes > 0 {
						someoneFlashed = true
						flashesThisStep += amountOfFlashes
					}
				}
			}

		}
		if flashesThisStep == 100 {
			fmt.Printf("step: %v", gameSteps+1)
			break
		}

		gameSteps++
	}
	fmt.Println("")
	fmt.Println(flashes)
}
