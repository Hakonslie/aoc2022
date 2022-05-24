package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const fileName = "crabs"

func main() {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	scanner.Scan()
	line := scanner.Text()
	crabsPositionString := strings.Split(line, ",")
	var crabsPosition []int
	for _, l := range crabsPositionString {
		ok, _ := strconv.Atoi(l)
		crabsPosition = append(crabsPosition, ok)
	}

	fuelCost := make(map[int]int)

	for v := 0; v < 1000; v++ {
		// already checked this row
		if fuelCost[v] != 0 {
			continue
		}

		for _, o := range crabsPosition {
			if o > v {
				for i := 1; i <= (o - v); i++ {
					fuelCost[v] += i
				}
			} else if o < v {
				for i := 1; i <= (v - o); i++ {
					fuelCost[v] += i
				}
			}
		}
	}

	mEff := 5000000000000
	for _, v := range fuelCost {
		if v < mEff {
			mEff = v
		}
	}
	fmt.Println(mEff)

}
