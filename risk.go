package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const fileName = "risk"

type node struct {
	connections []string
	risk        int
}

func main() {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var s [100][100]node
	row := 0
	nodes := make(map[string]node)
	for scanner.Scan() {
		for i, k := range strings.Split(scanner.Text(), "") {
			in, _ := strconv.Atoi(k)
			var c []string
			if i%1 == 0 {
				if row != 99 {
					c = append(c, fmt.Sprintf("%d,%d", row+1, i))
				}
				if i != 99 {
					c = append(c, fmt.Sprintf("%d,%d", row, i+1))
				}
			} else {
				if i != 99 {
					c = append(c, fmt.Sprintf("%d,%d", row, i+1))
				}
				if row != 99 {
					c = append(c, fmt.Sprintf("%d,%d", row+1, i))
				}
			}
			n := node{risk: in, connections: c}
			s[row][i] = n
		}
		row++
	}
	for i, l := range s {
		for j, o := range l {
			nodes[fmt.Sprintf("%d,%d", i, j)] = o
		}
	}

	context := make(map[string]int)
	context["best"] = 400
	context["current"] = 0

	search("1,0", "0,0", context, nodes)
	search("0,1", "0,0", context, nodes)

	fmt.Println(context["best"])
}

func search(c, p string, ctx map[string]int, nodes map[string]node) {
	crn := ctx["current"]
	ctx["current"] = crn + nodes[c].risk
	fmt.Println(c)

	if ctx["current"] >= ctx["best"] {
		fmt.Println("Too expensive")
		ctx["current"] = crn
		return
	}
	if c == "99,99" {
		ctx["best"] = ctx["current"]
		ctx["current"] = crn
		return
	}
	for _, v := range nodes[c].connections {
		if v == p {
			continue
		}
		search(v, c, ctx, nodes)
	}
	ctx["current"] = crn
}
