package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func firstTask() int {
	file, err := os.Open("depths")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	depth := 0
	depthIncreases := -1
	for scanner.Scan() {
		newDepth, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		if newDepth > depth {
			depthIncreases++
		}
		depth = newDepth
	}

	return depthIncreases
}
func secondTask() int {
	file, err := os.Open("depths")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	a := 0
	b := 0
	c := 0
	pos := 0
	depthIncreases := 0

	for scanner.Scan() {
		newDepth, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		total := a + b + c
		switch pos % 3 {
		case 0:
			a = newDepth
		case 1:
			b = newDepth
		case 2:
			c = newDepth
		}
		if pos > 2 {
			if a+b+c > total {
				depthIncreases++
				fmt.Println("increased")
			} else if a+b+c == total {
				fmt.Println("no change")
			} else {
				fmt.Println("decreased")
			}
		}
		pos++
	}
	return depthIncreases
}

func fourthTask() int {
	file, err := os.Open("move")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	horPos := 0
	verPos := 0
	aim := 0

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
		splitLine := strings.Split(line, " ")
		distance, _ := strconv.Atoi(splitLine[1])
		switch splitLine[0] {
		case "forward":
			horPos += distance
			verPos += aim * distance
		case "up":
			aim -= distance
		case "down":
			aim += distance
		}
	}
	return horPos * verPos

}

func readMostCommonBit(pos int, vals []string) string {

	positive := 0
	negative := 0

	for _, o := range vals {
		val := strings.Split(o, "")
		if val[pos] == strconv.Itoa(0) {
			negative++
		} else {
			positive++
		}
	}

	if positive > negative {
		str1 := strconv.Itoa(1)
		return str1
	} else {
		str0 := strconv.Itoa(0)
		return str0
	}
}

func fifthTask() int64 {
	vals := ""

	for i := 0; i < 12; i++ {
		file, err := os.Open("example")
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
		newVal := ""
		vals = vals + "" + newVal
	}
	valsFlipped := ""
	for _, o := range vals {
		if o == '1' {
			valsFlipped = valsFlipped + "" + "0"
		} else {
			valsFlipped = valsFlipped + "" + "1"
		}

	}

	i, _ := strconv.ParseInt(vals, 2, 64)
	o, _ := strconv.ParseInt(valsFlipped, 2, 64)
	return i * o
}
func sixthTask() int64 {
	file, err := os.Open("energy")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var input []string
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}
	firstInput := input
	secondInput := input
	return mynewfunc(firstInput, true) * mynewfunc(secondInput, false)
}
func mynewfunc(input []string, over bool) int64 {
	buildString := ""
	for i := 0; i < 12; i++ {
		if len(input) == 1 {
			continue
		}
		var falseAtPos int
		var trueAtPos int
		for _, j := range input {
			stn := strings.Split(j, "")
			if stn[i] == "0" {
				falseAtPos++
			} else {
				trueAtPos++
			}
		}
		fmt.Printf("False: %d, True: %d \n", falseAtPos, trueAtPos)
		trigger := ""
		if over {
			if trueAtPos >= falseAtPos {
				trigger = "1"
			} else {
				trigger = "0"
			}
		} else {
			if falseAtPos <= trueAtPos {
				trigger = "0"
			} else {
				trigger = "1"
			}
		}
		fmt.Println(trigger)

		buildString = buildString + trigger
		var dirtyInput []string
		for _, k := range input {
			if strings.Split(k, "")[i] == trigger {
				dirtyInput = append(dirtyInput, k)
			}
		}
		input = dirtyInput
	}

	nnbr, err := strconv.ParseInt(input[0], 2, 64)
	if err != nil {
		log.Fatal(err)
	}
	return nnbr
}

func main() {
	fmt.Println(sixthTask())
}
