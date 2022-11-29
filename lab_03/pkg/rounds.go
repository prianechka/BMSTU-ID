package desAlg

import (
	"log"
	"strconv"
	"strings"
)

func extendedPermutationE(s []string) []string {
	return []string{s[31], s[0], s[1], s[2], s[3], s[4], s[3], s[4],
		s[5], s[6], s[7], s[8], s[7], s[8], s[9], s[10],
		s[11], s[12], s[11], s[12], s[13], s[14], s[15], s[16],
		s[15], s[16], s[17], s[18], s[19], s[20], s[19], s[20],
		s[21], s[22], s[23], s[24], s[23], s[24], s[25], s[26],
		s[27], s[28], s[27], s[28], s[29], s[30], s[31], s[0]}
}

func finishedPermutationP(s []string) []string {
	return []string{s[15], s[6], s[19], s[20], s[28], s[11], s[27], s[16],
		s[0], s[14], s[22], s[25], s[4], s[17], s[30], s[9],
		s[1], s[7], s[23], s[13], s[31], s[26], s[2], s[8],
		s[18], s[12], s[29], s[5], s[21], s[10], s[3], s[24]}
}

func AddNulls(tmpString string, neededLength int) string {
	for len(tmpString) < neededLength {
		tmpString = "0" + tmpString
	}
	return tmpString
}

func EncryptRounds(bytes []string, keys []string) (l16 []string, r16 []string) {
	leftBlock := bytes[:32]
	rightBlock := bytes[32:]

	for i := 0; i < 16; i++ {
		feistel := Feistel(rightBlock, keys[i])
		leftBlockInt, _ := strconv.ParseUint(strings.Join(leftBlock, ""), 2, 0)

		tmpResult := leftBlockInt ^ feistel
		tmpString := strconv.FormatUint(tmpResult, 2)
		tmpString = AddNulls(tmpString, 32)
		tmpValue := strings.Split(tmpString, "")

		leftBlock = rightBlock
		rightBlock = tmpValue
	}

	return leftBlock, rightBlock
}

func Feistel(block []string, key string) uint64 {
	blockExpanded := extendedPermutationE(block)
	blockExpandedInt, _ := strconv.ParseUint(strings.Join(blockExpanded, ""), 2, 0)

	keyInt, _ := strconv.ParseUint(key, 2, 0)

	ZInFeistel := blockExpandedInt ^ keyInt
	ZInFeistelString := strconv.FormatUint(ZInFeistel, 2)
	ZInFeistelString = AddNulls(ZInFeistelString, 48)

	ZArray := strings.Split(ZInFeistelString, "")
	Z := subtitution(ZArray)
	Z = finishedPermutationP(Z)

	result, _ := strconv.ParseUint(strings.Join(Z, ""), 2, 0)

	return result
}

func DecryptRounds(bytes []string, keys []string) (l16 []string, r16 []string) {
	leftBlock := bytes[:32]
	rightBlock := bytes[32:]

	for i := 15; i >= 0; i-- {
		feistel := Feistel(leftBlock, keys[i])
		rightBlockInt, _ := strconv.ParseUint(strings.Join(rightBlock, ""), 2, 0)
		tmpResult := rightBlockInt ^ feistel
		tmpString := strconv.FormatUint(tmpResult, 2)
		tmpString = AddNulls(tmpString, 32)
		tmpValue := strings.Split(tmpString, "")

		rightBlock = leftBlock
		leftBlock = tmpValue
	}

	return leftBlock, rightBlock
}

func subtitution(bytes []string) []string {
	var (
		dividedBlocks = [][]string{bytes[:6], bytes[6:12], bytes[12:18], bytes[18:24],
			bytes[24:30], bytes[30:36], bytes[36:42], bytes[42:]}
		resultString       string
		substitutionResult int
	)

	for i, dividedBlock := range dividedBlocks {
		xMatrix, err := strconv.ParseUint(strings.Join(dividedBlock[1:5], ""), 2, 0)
		if err != nil {
			log.Fatal(err)
		}
		yMatrix, err := strconv.ParseUint(strings.Join([]string{dividedBlock[0], dividedBlock[5]}, ""), 2, 0)
		if err != nil {
			log.Fatal(err)
		}

		substitutionResult = substitutionMatrix[i][xMatrix][yMatrix]

		if substitutionResult <= 1 {
			resultString += "000" + strconv.FormatUint(uint64(substitutionResult), 2)
		} else if substitutionResult <= 3 {
			resultString += "00" + strconv.FormatUint(uint64(substitutionResult), 2)
		} else if substitutionResult <= 7 {
			resultString += "0" + strconv.FormatUint(uint64(substitutionResult), 2)
		} else {
			resultString += strconv.FormatUint(uint64(substitutionResult), 2)
		}

	}

	return strings.Split(resultString, "")
}
