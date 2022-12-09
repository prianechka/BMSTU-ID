package huffman

import (
	"encoding/json"
	"fmt"
	"math/big"
	"strconv"
)

func MakeArrayOfBytes(srcString []byte, chunkSize int) (resultChunks [][]byte) {
	if chunkSize >= len(srcString) {
		resultChunks = [][]byte{srcString}
	} else {
		tmpStringOfBytes := make([]byte, chunkSize)
		currentLength := 0
		for _, curByte := range srcString {
			tmpStringOfBytes[currentLength] = curByte
			currentLength++
			if currentLength == chunkSize {
				resultChunks = append(resultChunks, tmpStringOfBytes)
				currentLength = 0
				tmpStringOfBytes = make([]byte, chunkSize)
			}
		}
		if currentLength > 0 {
			resultChunks = append(resultChunks, tmpStringOfBytes[:currentLength])
		}
	}
	return resultChunks
}

func CreateBytesFromBits(chunks [][]byte) []byte {
	result := make([]byte, 0)
	for _, chunk := range chunks {
		chunkString := fmt.Sprintf("%s", chunk)
		resultInt, _ := strconv.ParseUint(chunkString, 2, 0)
		result = append(result, byte(resultInt))
	}
	return result
}

func ReadFileWithLen(data []byte, length int) []byte {
	return data[:length]
}

func WriteTree(tree *Tree) ([]byte, error) {
	return json.Marshal(tree)
}

func CreateTreeFromFile(data []byte) (result *Tree) {
	_ = json.Unmarshal(data, &result)
	return result
}

func ConvertFromFileBytesToMyBytes(data []byte) []byte {
	result := make([]byte, 0)
	for _, b := range data[:len(data)-1] {
		number := big.NewInt(int64(b)).Text(2)
		for len(number) < 8 {
			number = fmt.Sprintf("0%s", number)
		}
		for _, c := range number {
			resultInt, _ := strconv.ParseUint(string(c), 10, 0)
			result = append(result, byte(resultInt))
		}
	}
	lastByte := data[len(data)-1]
	number := big.NewInt(int64(lastByte)).Text(2)
	for _, c := range number {
		resultInt, _ := strconv.ParseUint(string(c), 10, 0)
		result = append(result, byte(resultInt))
	}
	return result
}
