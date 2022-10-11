package enigma

type Machine struct {
	Rotors    []*MyRotor
	Reflector *MyReflector
}

func InitMachine(seed int64) *Machine {
	result := Machine{
		Rotors:    []*MyRotor{MakeRotor(), MakeRotor(), MakeRotor()},
		Reflector: &MyReflector{},
	}
	for _, el := range result.Rotors {
		el.FillRotor(seed)
	}
	result.Reflector.MakeAlfabet()
	return &result
}
func (enigma *Machine) encodeByte(data byte) byte {
	for _, curRotor := range enigma.Rotors {
		data = curRotor.GetByteStraight(data)
	}
	data = enigma.Reflector.Bytes[data]
	for i, _ := range enigma.Rotors {
		data = enigma.Rotors[len(enigma.Rotors)-i-1].GetByteReverse(data)
	}

	for _, rotor := range enigma.Rotors {
		rotor.RotateRotor()
		if rotor.CurRotates > 0 {
			break
		}
	}

	return data
}

func (enigma *Machine) EncodeText(data []byte) []byte {
	resultText := make([]byte, 0)
	for _, b := range data {
		resultText = append(resultText, enigma.encodeByte(b))
	}
	return resultText
}
