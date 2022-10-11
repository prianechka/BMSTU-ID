package enigma

import "math"

const AlfabetLength = math.MaxUint8 + 1

type Reflector interface {
	Reflect([]byte) []byte
}

type MyReflector struct {
	Bytes []byte
}

func (reflect *MyReflector) MakeAlfabet() {
	reflect.Bytes = make([]byte, AlfabetLength)
	for i := 0; i < AlfabetLength; i++ {
		reflect.Bytes[i] = byte(AlfabetLength - i)
	}
}
