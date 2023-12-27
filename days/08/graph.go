package main

import (
	"errors"
	"fmt"
)

type Node struct {
	Id string
	Left *Node
	Right *Node
}

type Graph struct {
	Nodes map[string]*Node
}


func NewGraph() *Graph {
	return &Graph{
		Nodes: make(map[string]*Node),
	}
}

func (g *Graph) AddNode(Id string) {
	node := &Node{Id: Id}
	g.Nodes[Id] = node
}

func (g *Graph) AddEdge(from, to string, left bool) error {
	fromNode, fromOk := g.Nodes[from]
	toNode, toOk := g.Nodes[to]
	if !fromOk {
		return errors.New(fmt.Sprintf("Unable to add edge, no Node with Id %s exists!", from))
	}
	if !toOk {
		return errors.New(fmt.Sprintf("Unable to add edge, no Node with Id %s exists!", to))
	}
	if left {
		fromNode.Left = toNode
	} else {
		fromNode.Right = toNode
	}
	return nil
}

func (g *Graph) CountFindSteps(start, instructions string, loopCondition func(*Node) bool) (int, error) {
	startNode, startOk := g.Nodes[start]
	if !startOk {
		return -1, errors.New("Start node does not exist in graph!")
	}
	currentNode := startNode
	step := 0
	for loopCondition(currentNode) {
		instruction := string(instructions[step % len(instructions)])
		if instruction == "L" {
			currentNode = currentNode.Left
		} else {
			currentNode = currentNode.Right
		}
		step += 1
	}
	return step, nil
}