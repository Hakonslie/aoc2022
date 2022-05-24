package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const fileName = "polymer"

func main() {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Scan()

	template := strings.Split(scanner.Text(), "")
	pairs := make(map[string]int)
	counter := make(map[string]int)

	counter[template[0]] = counter[template[0]] + 1
	for i := 0; i < len(template)-1; i++ {
		counter[template[i+1]] = counter[template[i+1]] + 1
		str := template[i] + template[i+1]
		pairs[str] = pairs[str] + 1
	}

	scanner.Scan()
	inRules := make(map[string]string)
	for scanner.Scan() {
		str := strings.Split(scanner.Text(), " -> ")
		inRules[str[0]] = str[1]
	}

	iterations := 40
	for i := 0; i < iterations; i++ {
		tempPairs := make(map[string]int)
		// Copy state
		for k, v := range pairs {
			tempPairs[k] = v
		}
		for k, v := range pairs {
			if v > 0 {
				newChar := inRules[k]
				if newChar != "" {
					// Split pair
					sp := strings.Split(k, "")
					tempPairs[k] = tempPairs[k] - v

					nOne := sp[0] + newChar
					nTwo := newChar + sp[1]
					tempPairs[nOne] = tempPairs[nOne] + v
					tempPairs[nTwo] = tempPairs[nTwo] + v

					// Add to counter
					counter[newChar] = counter[newChar] + v
				}
			}
		}
		// set new state
		pairs = tempPairs
	}

	largest := template[0]
	smallest := template[0]
	for i := 0; i < 50; i++ {
		for k, v := range counter {
			if v < counter[smallest] {
				smallest = k
			}
			if v > counter[largest] {
				largest = k
			}
		}
	}
	fmt.Println(counter[largest] - counter[smallest])
}
