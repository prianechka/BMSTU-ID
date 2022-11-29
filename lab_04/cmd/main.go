package main

import (
	"bytes"
	"errors"
	"fmt"
	alg "lab_04/pkg"
	"log"
	"os"
)

const (
	Limit         = 2048
	AllPrivileges = 0777
)

var (
	inputFile, outputFile, checkFile string
)

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

func main() {
	err := CheckArgs()
	if err == nil {
		data, _ := os.ReadFile(inputFile)
		newRSA := alg.NewRSA(Limit)
		encryptedInfo := newRSA.Encrypt(data)
		encryptedBytes := alg.Int64ToBytes(encryptedInfo)

		err = os.WriteFile(outputFile, encryptedBytes, AllPrivileges)
		if err != nil {
			log.Fatalf("Can't write encrypted data, error is: %s", err)
		}

		decrypted := newRSA.Decrypt(encryptedInfo)
		_ = os.WriteFile(checkFile, decrypted, AllPrivileges)

		isEqual := bytes.Equal(data, decrypted)
		if isEqual {
			fmt.Printf("Input and decoded files are equal: %t\n", isEqual)
		} else {
			fmt.Printf("Bad result: %t\n", isEqual)
		}
	}
}
