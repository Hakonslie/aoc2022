package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const fileName = "lava"

func searcher(row, col int, depthMap [][]int, searches map[string]bool) int {
	coord := fmt.Sprintf("%d,%d", row, col)
	if searches[coord] || row < 0 || row >= len(depthMap) || col < 0 || col >= len(depthMap[row]) {
		return 0
	} else if depthMap[row][col] == 9 {
		return 0
	}

	searches[coord] = true
	n := searcher(row-1, col, depthMap, searches)
	w := searcher(row, col-1, depthMap, searches)
	s := searcher(row+1, col, depthMap, searches)
	e := searcher(row, col+1, depthMap, searches)
	return n + w + s + e + 1
}

func main() {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var depthMap [][]int
	for scanner.Scan() {
		var row []int
		for _, v := range strings.Split(scanner.Text(), "") {
			n, _ := strconv.Atoi(v)
			row = append(row, n)
		}
		depthMap = append(depthMap, row)
	}
	/*
		riskLevel := 0
		for i := 0; i < len(depthMap); i++ {
			for j := 0; j < len(depthMap[i]); j++ {
				currentDepth := depthMap[i][j]
				// inbounds NORTH
				if i-1 >= 0 {
					if currentDepth >= depthMap[i-1][j] {
						continue
					}
				}
				// inbounds WEST
				if j-1 >= 0 {
					if currentDepth >= depthMap[i][j-1] {
						continue
					}
				}
				// inbounds SOUTH
				if i+1 < len(depthMap) {
					if currentDepth >= depthMap[i+1][j] {
						continue
					}
				}
				// inbounds EAST
				if j+1 < len(depthMap[i]) {
					if currentDepth >= depthMap[i][j+1] {
						continue
					}
				}
				riskLevel += currentDepth + 1
			}
		}
	*/
	searches := make(map[string]bool)

	basins := [3]int{0, 0, 0}
	for i := 0; i < len(depthMap); i++ {
		for j := 0; j < len(depthMap[i]); j++ {
			result := searcher(i, j, depthMap, searches)
			if result > 0 {
				smallestBasin := 0
				for k, v := range basins {
					if v < basins[smallestBasin] {
						smallestBasin = k
					}
				}
				if result > basins[smallestBasin] {
					basins[smallestBasin] = result
				}
				fmt.Println(basins)
			}
		}
	}
	result := 1
	for _, v := range basins {
		result = result * v
	}

	fmt.Println(result)
}
