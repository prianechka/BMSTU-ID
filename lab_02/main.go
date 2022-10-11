package main

import (
	"bytes"
	"errors"
	"fmt"
	"lab_02/pkg"
	"log"
	"os"
	"time"
)

var (
	inputFile, outputFile, checkFile string
)

func main() {
	err := CheckArgs()
	if err != nil {
		log.Fatalln(err)
	}
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
	err = os.WriteFile(outputFile, encodedText, 0777)
	if err != nil {
		log.Fatalln(err)
	}

	enigmaMachine.Rotors = rotorsCopy
	decodedText := enigmaMachine.EncodeText(encodedText)
	err = os.WriteFile(checkFile, decodedText, 0777)
	if err != nil {
		log.Fatalln(err)
	}

	isEqual := bytes.Equal(data, decodedText)
	fmt.Printf("Input and decoded files are equal: %t\n", isEqual)
}

func CheckArgs() (err error) {
	if len(os.Args) != 4 {
		err = errors.New("bad args")
	} else {
		inputFile = os.Args[1]
		outputFile = os.Args[2]
		checkFile = os.Args[3]
	}
	return err
}
