package huffman

import (
	"math"
)

func NewTable(srcFileData []byte) *Table {
	table := make(map[uint32]*Node)
	for _, b := range srcFileData {
		if _, ok := table[uint32(b)]; !ok {
			table[uint32(b)] = &Node{value: uint32(b), weight: 1}
		} else {
			table[uint32(b)].weight++
		}
	}
	return &Table{Table: table}
}

func JoinNodes(firstNode, secondNode *Node, value uint32) *Node {
	return &Node{
		Left:   firstNode,
		Right:  secondNode,
		weight: firstNode.weight + secondNode.weight,
		value:  value,
	}
}

type Table struct {
	Table map[uint32]*Node
}

func (hft *Table) GetSmallestNode() *Node {
	smallest := &Node{weight: math.MaxInt32}
	for _, v := range hft.Table {
		if v.weight < smallest.weight {
			smallest = v
		}
	}
	return smallest
}
