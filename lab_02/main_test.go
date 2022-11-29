package main

import (
	"bytes"
	enigma "lab_02/pkg"
	"os"
	"testing"
	"time"
)

func Test(t *testing.T) {
	testFiles := []string{"tests/input.txt", "tests/image.jpg", "tests/archive.zip",
		"tests/im.png", "tests/main.out", "tests/test.o"}

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
		} else {
			t.Logf("%s passed", testFile)
		}
	}
}
