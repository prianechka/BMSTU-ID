package desAlg

import (
	"fmt"
)

func TranslateByteToBinaryFormat(inputString []byte) string {
	var binaryString string
	for _, symbol := range inputString {
		binaryString = fmt.Sprintf("%s%.8b", binaryString, symbol)
	}
	return binaryString
}

func CompleteKey(key []byte) []byte {
	for len(key)%8 > 0 {
		key = append(key, byte(0))
	}
	return key[:8]
}
