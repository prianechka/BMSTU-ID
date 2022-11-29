package desAlg

import (
	"strconv"
	"strings"
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

func AddPointToCompleteString(chunk []byte, length int) []byte {
	for len(chunk) < length {
		chunk = append(chunk, byte(0))
	}
	return chunk
}

func ConvertToBytes(result string) []byte {
	var length = len(result)
	var resultBytes = make([]byte, 0)
	for i := 0; i < length; i += 8 {
		tmpString := result[i : i+8]
		tmpByteValue, _ := strconv.ParseInt(tmpString, 2, 64)
		resultBytes = append(resultBytes, byte(tmpByteValue))
	}
	return resultBytes
}

func Encrypt(data []byte, keys []string) []byte {
	var result string
	chunks := MakeArrayOfBytes(data, 8)

	for _, chunk := range chunks {
		chunk = AddPointToCompleteString(chunk, 8)
		binaryMessage := TranslateByteToBinaryFormat(chunk)
		binarySlice := strings.Split(binaryMessage, "")
		binarySliceAfterIP := initPermutationIP(binarySlice)

		l16, r16 := EncryptRounds(binarySliceAfterIP, keys)
		lr16 := append(l16, r16...)
		result += strings.Join(finishPermutationIP(lr16), "")
	}

	return ConvertToBytes(result)
}

func Decrypt(data []byte, keys []string) []byte {
	var result string
	chunks := MakeArrayOfBytes(data, 8)

	for _, chunk := range chunks {
		chunk = AddPointToCompleteString(chunk, 8)
		binaryMessage := TranslateByteToBinaryFormat(chunk)
		binarySlice := strings.Split(binaryMessage, "")
		binarySliceAfterIP := initPermutationIP(binarySlice)

		l16, r16 := DecryptRounds(binarySliceAfterIP, keys)
		lr16 := append(l16, r16...)
		result += strings.Join(finishPermutationIP(lr16), "")
	}
	return ConvertToBytes(result)
}

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func initPermutationIP(s []string) []string {
	return []string{s[57], s[49], s[41], s[33], s[25], s[17], s[9], s[1],
		s[59], s[51], s[43], s[35], s[27], s[19], s[11], s[3],
		s[61], s[53], s[45], s[37], s[29], s[21], s[13], s[5],
		s[63], s[55], s[47], s[39], s[31], s[23], s[15], s[7],
		s[56], s[48], s[40], s[32], s[24], s[16], s[8], s[0],
		s[58], s[50], s[42], s[34], s[26], s[18], s[10], s[2],
		s[60], s[52], s[44], s[36], s[28], s[20], s[12], s[4],
		s[62], s[54], s[46], s[38], s[30], s[22], s[14], s[6]}
}

func finishPermutationIP(s []string) []string {
	return []string{s[39], s[7], s[47], s[15], s[55], s[23], s[63], s[31],
		s[38], s[6], s[46], s[14], s[54], s[22], s[62], s[30],
		s[37], s[5], s[45], s[13], s[53], s[21], s[61], s[29],
		s[36], s[4], s[44], s[12], s[52], s[20], s[60], s[28],
		s[35], s[3], s[43], s[11], s[51], s[19], s[59], s[27],
		s[34], s[2], s[42], s[10], s[50], s[18], s[58], s[26],
		s[33], s[1], s[41], s[9], s[49], s[17], s[57], s[25],
		s[32], s[0], s[40], s[8], s[48], s[16], s[56], s[24]}
}
