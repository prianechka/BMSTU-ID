package huffman

import (
	"errors"
	"math"
)

func NewHuffmanTree(srcFileData []byte) *Tree {
	huffmanTable := NewTable(srcFileData)
	var j *Node
	var i uint32 = math.MaxUint8 + 1

	for len(huffmanTable.Table) > 1 {
		firstSmallestNode := huffmanTable.GetSmallestNode()
		delete(huffmanTable.Table, firstSmallestNode.value)

		secondSmallestNode := huffmanTable.GetSmallestNode()
		delete(huffmanTable.Table, secondSmallestNode.value)

		j = JoinNodes(firstSmallestNode, secondSmallestNode, i)
		i++
		huffmanTable.Table[j.value] = j
	}
	return &Tree{Root: j, Data: srcFileData}
}

type Node struct {
	value  uint32 `json:"Value"`
	weight int    `json:"Weight"`
	Left   *Node  `json:"Left"`
	Right  *Node  `json:"Right"`
}

type Tree struct {
	Data []byte `json:"Data"`
	Root *Node  `json:"Root"`
}

func (t *Tree) Compress(res *[]byte) error {
	if t.Root == nil {
		return errors.New("root cannot be null")
	}
	for _, c := range t.Data {
		stack := Stack{}
		t.encodeByte(c, stack.New(), t.Root, res)
	}
	return nil
}

func (t *Tree) Decompress(encoded []byte) ([]byte, error) {
	if t.Root == nil {
		return nil, errors.New("root cannot be null")
	}

	n := t.Root
	var decoded []byte
	for _, b := range encoded {
		if b == '0' {
			n = n.Left
		}
		if b == '1' {
			n = n.Right
		}
		if n.Left == nil && n.Right == nil {
			decoded = append(decoded, byte(n.value))
			n = t.Root
		}
	}

	return decoded, nil
}

func (t *Tree) encodeByte(curSymbol byte, s *Stack, rootNode *Node, resultArrayOfBytes *[]byte) {
	if rootNode.Left == nil && rootNode.Right == nil {
		if rootNode.value == uint32(curSymbol) {
			for _, i := range s.items {
				*resultArrayOfBytes = append(*resultArrayOfBytes, byte(i))
			}
		}
		return
	}
	s.Push('0')
	t.encodeByte(curSymbol, s, rootNode.Left, resultArrayOfBytes)
	s.Pop()

	s.Push('1')
	t.encodeByte(curSymbol, s, rootNode.Right, resultArrayOfBytes)
	s.Pop()
}
