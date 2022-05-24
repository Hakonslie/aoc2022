package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const fileName = "displays"

func main() {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var decoder [][]string
	var outputs [][]string

	for scanner.Scan() {
		out := strings.Split(scanner.Text(), " | ")
		decoder = append(decoder, strings.Split(out[0], " "))
		outputs = append(outputs, strings.Split(out[1], " "))
	}

	var numbersArray []int

	for i, v := range decoder {
		// each index represents a position from left to right, top to bottom
		var decoding [7]string

		inZero := make(map[string]bool)
		inOne := make(map[string]bool)
		inTwo := make(map[string]bool)
		inThree := make(map[string]bool)
		inFour := make(map[string]bool)
		inSix := make(map[string]bool)
		inSeven := make(map[string]bool)
		inEight := make(map[string]bool)
		inNine := make(map[string]bool)
		// range over all decoding inputs (strings)
		for _, o := range v {
			switch len(o) {
			case 2:
				for _, j := range strings.Split(o, "") {
					inOne[j] = true
				}
			case 3:
				for _, j := range strings.Split(o, "") {
					inSeven[j] = true
				}
			case 4:
				for _, j := range strings.Split(o, "") {
					inFour[j] = true
				}
			case 7:
				for _, j := range strings.Split(o, "") {
					inEight[j] = true
				}
			}
		}
		// Number seven shares 2 lines with number one, so last on can be deducted
		for k, _ := range inSeven {
			if !inOne[k] {
				decoding[0] = k
				break
			}
		}

		// range all decoding inputs again, ( do it twice to get all )
		for i := 0; i < 3; i++ {
			for _, o := range v {
				if len(o) == 5 {
					hasInOne := make(map[string]bool)
					for _, l := range strings.Split(o, "") {
						if inOne[l] {
							hasInOne[l] = true
						}
					}
					hasInFour := make(map[string]bool)
					for _, l := range strings.Split(o, "") {
						if inFour[l] {
							hasInFour[l] = true
						}
					}
					if len(hasInOne) == 2 {
						// 3
						for _, b := range strings.Split(o, "") {
							inThree[b] = true
						}
						for s, _ := range inNine {
							if !inThree[s] {
								decoding[1] = s
							}
						}
					} else if len(hasInFour) == 2 {
						// 2
						for _, b := range strings.Split(o, "") {
							inTwo[b] = true
						}
						for k, _ := range inEight {
							if !inTwo[k] && k != decoding[1] {
								decoding[5] = k
							}
						}

					}
				} else if len(o) == 6 {
					notInFour := make(map[string]bool)
					for _, l := range strings.Split(o, "") {
						if !inFour[l] {
							notInFour[l] = true
						}
					}
					hasInOne := make(map[string]bool)
					for _, l := range strings.Split(o, "") {
						if inOne[l] {
							hasInOne[l] = true
						}
					}
					if len(notInFour) > 2 {
						if len(hasInOne) != 2 {
							// 6
							for _, l := range strings.Split(o, "") {
								inSix[l] = true
							}
							for k, _ := range inEight {
								if !inSix[k] {
									decoding[2] = k
								}
							}
						} else {
							// 0
							for _, l := range strings.Split(o, "") {
								inZero[l] = true
							}
							for k, _ := range inEight {
								if !inZero[k] {
									decoding[3] = k
								}
							}
						}

					} else {
						// 9
						if notInFour[decoding[0]] {
							for a, _ := range notInFour {
								if a != decoding[0] {
									decoding[6] = a
								}
							}
						}
						for _, j := range strings.Split(o, "") {
							inNine[j] = true
						}
					}
				}
			}
		}

		for k, _ := range inEight {
			if !inNine[k] {
				decoding[4] = k
			}
		}

		// Map decoding

		decodingMap := make(map[string]int)
		for j, v := range decoding {
			decodingMap[v] = j
		}

		numberBuilder := ""
		for _, v := range outputs[i] {
			intBuilder := make(map[int]bool)
			for _, p := range strings.Split(v, "") {
				intBuilder[decodingMap[p]] = true
			}
			// Time to build the numbers
			switch len(intBuilder) {
			case 2:
				numberBuilder = numberBuilder + "1"
			case 3:
				numberBuilder = numberBuilder + "7"
			case 4:
				numberBuilder = numberBuilder + "4"
			case 5:
				// 2, 3, 5
				hasTwo := false
				hasFive := false
				for k, _ := range intBuilder {
					if k == 2 {
						hasTwo = true
					} else if k == 5 {
						hasFive = true
					}
				}
				if hasFive && hasTwo {
					numberBuilder = numberBuilder + "3"
				} else if hasFive && !hasTwo {
					numberBuilder = numberBuilder + "5"
				} else if !hasFive && hasTwo {
					numberBuilder = numberBuilder + "2"
				}
			case 6:
				// 0, 6, 9
				hasThree := false
				hasTwo := false
				for k, _ := range intBuilder {
					if k == 3 {
						hasThree = true
					} else if k == 2 {
						hasTwo = true
					}
				}
				if !hasThree {
					numberBuilder = numberBuilder + "0"
				} else if !hasTwo {
					numberBuilder = numberBuilder + "6"
				} else {
					numberBuilder = numberBuilder + "9"
				}
			case 7:
				numberBuilder = numberBuilder + "8"
			}
		}
		numb, _ := strconv.Atoi(numberBuilder)
		numbersArray = append(numbersArray, numb)
	}
	total := 0
	for _, v := range numbersArray {
		total += v
	}
	fmt.Println(total)
}
