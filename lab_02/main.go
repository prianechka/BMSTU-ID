package main

import (
	"bytes"
	"fmt"
	"lab_02/pkg"
	"log"
	"os"
	"time"
)

const (
	inputFile  = "input.txt"
	outputFile = "output.txt"
	checkFile  = "check.txt"
)

func main() {
	data, err := os.ReadFile(inputFile)
	if err != nil {
		log.Fatalln(err)
	}
	seed := time.Now().UnixNano()
	enigmaMachine := enigma.InitMachine(seed)
	rotorsCopy := []*enigma.MyRotor{enigma.MakeRotor(), enigma.MakeRotor(), enigma.MakeRotor()}
	for _, el := range rotorsCopy {
		el.FillRotor(seed)
	}

	encodedText := enigmaMachine.EncodeText(data)
	err = os.WriteFile(outputFile, encodedText, 777)
	if err != nil {
		log.Fatalln(err)
	}

	enigmaMachine.Rotors = rotorsCopy
	decodedText := enigmaMachine.EncodeText(encodedText)
	err = os.WriteFile(checkFile, decodedText, 777)
	if err != nil {
		log.Fatalln(err)
	}

	isEqual := bytes.Equal(data, decodedText)
	fmt.Printf("Input and decoded files are equal: %t\n", isEqual)
}
