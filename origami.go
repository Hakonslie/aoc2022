package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const fileName = "origami"

type coord struct {
	x, y int
}
type fold struct {
	x bool
	i int
}

func main() {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var coords []coord
	var folds []fold
	for scanner.Scan() {
		txt := scanner.Text()
		matched, _ := regexp.Match(`fold`, []byte(txt))
		if matched {
			ok := strings.Split(strings.Split(txt, " ")[2], "=")
			iInt, _ := strconv.Atoi(ok[1])
			folds = append(folds, fold{x: (ok[0] == "x"), i: iInt})
		} else if txt != "" {
			splitter := strings.Split(txt, ",")
			xi, _ := strconv.Atoi(splitter[0])
			yi, _ := strconv.Atoi(splitter[1])
			coords = append(coords, coord{x: xi, y: yi})
		}
	}
	for _, f := range folds {
		for i, c := range coords {
			if !f.x && c.y > f.i {
				coords[i].y = c.y - ((c.y - f.i) * 2)
			}
			if f.x && c.x > f.i {
				coords[i].x = c.x - ((c.x - f.i) * 2)
			}
		}
	}
	mC := make(map[string]bool)
	for _, v := range coords {
		mC[fmt.Sprintf("%d,%d", v.x, v.y)] = true
	}
	var written [10][40]bool
	for i, o := range written {
		for j, _ := range o {
			if mC[fmt.Sprintf("%d,%d", j, i)] {
				fmt.Printf("#")
			} else {
				fmt.Printf(" ")
			}
		}
		fmt.Println("")
	}
}
