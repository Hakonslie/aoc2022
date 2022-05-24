package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const fileName = "path"

type node struct {
	linked []string
	large  bool
}

func traverse(nod string, context map[int]bool, visited map[string]int, nodes map[string]node) int {

	if nod == "end" {
		return 1
	}
	visitedThisSmallCave := false
	if visited[nod] > 0 && !nodes[nod].large {
		if context[0] == false {
			context[0] = true
			visitedThisSmallCave = true
		} else {
			return 0
		}
	}

	visited[nod] = visited[nod] + 1

	// Must. Go. Further!
	paths := 0
	for _, v := range nodes[nod].linked {
		// Dont go back to start
		if v != "start" {
			paths += traverse(v, context, visited, nodes)
		}
	}
	// OMG IS THIS IT ?!?!
	visited[nod] = visited[nod] - 1
	if visitedThisSmallCave {
		context[0] = false
	}

	return paths

}

func main() {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	nodes := make(map[string]node)
	// Create nodes
	for _, v := range lines {
		for _, l := range strings.Split(v, "-") {
			var emptyNodes []string
			// Is large if matches with itself in uppercase
			large := strings.ToUpper(l) == l
			newNode := node{large: large, linked: emptyNodes}
			nodes[l] = newNode
		}
	}
	// Add Links
	for _, v := range lines {
		gutted := strings.Split(v, "-")

		// Add one direction
		tempNode := nodes[gutted[0]]
		tempNode.linked = append(tempNode.linked, gutted[1])
		nodes[gutted[0]] = tempNode

		// Add second direction
		tempNode = nodes[gutted[1]]
		tempNode.linked = append(tempNode.linked, gutted[0])
		nodes[gutted[1]] = tempNode
	}

	paz := 0
	visited := make(map[string]int)
	context := make(map[int]bool)
	context[0] = false // Small cave visited once
	paz += traverse("start", context, visited, nodes)
	fmt.Println(paz)
}
