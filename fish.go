package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const fileName = "fish"

func main() {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	scanner.Scan()
	line := scanner.Text()
	fishIntervalsString := strings.Split(line, ",")
	mapOfFish := make(map[int]int)
	for i := 0; i < 9; i++ {
		mapOfFish[i] = 0
	}
	for _, o := range fishIntervalsString {
		i, _ := strconv.Atoi(o)
		mapOfFish[i]++
	}

	intervalDays := 256
	for i := 0; i < intervalDays; i++ {
		temp := mapOfFish[0]
		mapOfFish[9] = mapOfFish[0]
		for i := 1; i < 10; i++ {
			mapOfFish[i-1] = mapOfFish[i]
		}
		mapOfFish[6] += temp
	}

	totalFish := 0
	mapOfFish[9] = 0
	for _, v := range mapOfFish {

		totalFish += v
	}
	fmt.Printf("After %d days: %d \n", intervalDays, totalFish)

}
