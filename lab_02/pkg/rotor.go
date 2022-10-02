package enigma

import (
	"math/rand"
)

type Rotor interface {
	FillRotor(seed int64)
	GetByteStraight(byte) byte
	GetByteReverse(byte) byte
	RotateRotor()
}

type MyRotor struct {
	Bytes      []byte
	CurRotates int
}

func MakeRotor() *MyRotor {
	return &MyRotor{
		Bytes:      []byte{},
		CurRotates: 0,
	}
}

func (rotor *MyRotor) FillRotor(seed int64) {
	rotor.Bytes = make([]byte, AlfabetLength)
	for i := 0; i < AlfabetLength; i++ {
		rotor.Bytes[i] = byte(i)
	}

	rand.Seed(seed)
	rand.Shuffle(len(rotor.Bytes), func(i, j int) { rotor.Bytes[i], rotor.Bytes[j] = rotor.Bytes[j], rotor.Bytes[i] })
}

func (rotor *MyRotor) GetByteStraight(text byte) byte {
	return rotor.Bytes[text]
}

func (rotor *MyRotor) GetByteReverse(text byte) (result byte) {
	for index, el := range rotor.Bytes {
		if el == text {
			result = byte(index)
		}
	}
	return result
}

func (rotor *MyRotor) RotateRotor() {
	rotor.Bytes = append([]byte{rotor.Bytes[AlfabetLength-1]}, rotor.Bytes[:AlfabetLength-1]...)
	rotor.CurRotates++
	if rotor.CurRotates == AlfabetLength {
		rotor.CurRotates = 0
	}
}
