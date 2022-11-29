package desAlg

import (
	"strings"
)

func GenerateKeys(key []byte) (keys []string) {
	initialKey := TranslateByteToBinaryFormat(key)
	initialKeySlice := strings.Split(initialKey, "")
	leftBlock, rightBlock := initPermutationB(initialKeySlice)

	shiftIndex := []int{1, 1, 2, 2, 2, 2, 2, 2, 1, 2, 2, 2, 2, 2, 2, 1}

	for i := 0; i < 16; i++ {
		leftBlock = leftShift(leftBlock, shiftIndex[i])
		rightBlock = leftShift(rightBlock, shiftIndex[i])
		roundKey := append(leftBlock, rightBlock...)
		keys = append(keys, strings.Join(compressPermutationCP(roundKey), ""))
	}

	return keys
}

func initPermutationB(s []string) ([]string, []string) {
	return []string{
			s[56], s[48], s[40], s[32], s[24], s[16], s[8],
			s[0], s[57], s[49], s[41], s[33], s[25], s[17],
			s[9], s[1], s[58], s[50], s[42], s[34], s[26],
			s[18], s[10], s[2], s[59], s[51], s[43], s[35]},
		[]string{
			s[62], s[54], s[46], s[38], s[30], s[22], s[14],
			s[6], s[61], s[53], s[45], s[37], s[29], s[21],
			s[13], s[5], s[60], s[52], s[44], s[36], s[28],
			s[20], s[12], s[4], s[27], s[19], s[11], s[3]}
}

func compressPermutationCP(s []string) []string {
	return []string{
		s[13], s[16], s[10], s[23], s[0], s[4], s[2], s[27],
		s[14], s[5], s[20], s[9], s[22], s[18], s[11], s[3],
		s[25], s[7], s[15], s[6], s[26], s[19], s[12], s[1],
		s[40], s[51], s[30], s[36], s[46], s[54], s[29], s[39],
		s[50], s[44], s[32], s[47], s[43], s[48], s[38], s[55],
		s[33], s[52], s[45], s[41], s[49], s[35], s[28], s[31]}
}

func leftShift(bytes []string, shiftValue int) []string {
	switch shiftValue {
	case 1:
		first := bytes[0]
		for curByte := range bytes {
			if curByte == len(bytes)-1 {
				bytes[curByte] = first
			} else {
				bytes[curByte] = bytes[curByte+1]
			}
		}
	case 2:
		firsts := []string{bytes[0], bytes[1]}
		for curByte := range bytes {
			if curByte == len(bytes)-2 {
				bytes[curByte] = firsts[0]
			} else if curByte == len(bytes)-1 {
				bytes[curByte] = firsts[1]
			} else {
				bytes[curByte] = bytes[curByte+2]
			}
		}
	}
	return bytes
}
