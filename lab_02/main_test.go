package main

import (
	"bytes"
	enigma "lab_02/pkg"
	"os"
	"testing"
	"time"
)

func Test(t *testing.T) {
	testFiles := []string{"input.txt", "image.jpg", "im.png", "main.out", "test.o"}

	for _, testFile := range testFiles {
		data, _ := os.ReadFile(testFile)
		seed := time.Now().UnixNano()
		enigmaMachine := enigma.InitMachine(seed)
		rotorsCopy := []*enigma.MyRotor{enigma.MakeRotor(), enigma.MakeRotor(), enigma.MakeRotor()}
		for _, el := range rotorsCopy {
			el.FillRotor(seed)
		}
		encodedText := enigmaMachine.EncodeText(data)
		enigmaMachine.Rotors = rotorsCopy
		decodedText := enigmaMachine.EncodeText(encodedText)

		isEqual := bytes.Equal(data, decodedText)
		if !isEqual {
			t.Errorf("not equal")
		}
	}
}
