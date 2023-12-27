package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

// Parse input into a graph structure
func Parse(path string) (string, *Graph){
	content, _ := os.ReadFile(path)
	lines := strings.Split(string(content), "\n")
	instructions := lines[0]
	network := NewGraph()
	nodeExp := regexp.MustCompile(`[A-Z\d]{3}`)

	var nodes []string
	var edges [][]string
	for _, line := range lines[2:] {
		result := nodeExp.FindAllString(line, -1)
		nodes = append(nodes, result[0])
		network.AddNode(result[0])
		edges = append(edges, result[1:])
	}
	// Now all the nodes are defined and the edges can be added:
	for i, edge := range edges {
		network.AddEdge(nodes[i], edge[0], true)
		network.AddEdge(nodes[i], edge[1], false)
	}
	return instructions, network
}

func LoopCondiditionPart1(node *Node) bool {
	return node.Id != "ZZZ"
}

func LoopCondiditionPart2(node *Node) bool {
	return !strings.HasSuffix(node.Id, "Z")
}



func SolvePart2(graph Graph, instructions string) int{
	var startNodes []string
	for _, node := range(graph.Nodes) {
		if strings.HasSuffix(node.Id, "A") {
			startNodes = append(startNodes, node.Id)
		}
	}
	var steps []int
	for _, startNode := range startNodes {
		step, _ := graph.CountFindSteps(startNode, instructions, LoopCondiditionPart2)
		steps = append(steps, step)
	}
	return lcmSlice(steps)
}

func main() {
	instructions, network := Parse("input.txt")
	solution, _ := network.CountFindSteps("AAA", instructions, LoopCondiditionPart1)
	solution2 := SolvePart2(*network, instructions)
	fmt.Printf("Solution part 1: %d\n", solution)
	fmt.Printf("Solution part 2: %d\n", solution2)
}