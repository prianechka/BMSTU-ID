package main

import (
	"bytes"
	"desAlg/pkg"
	"errors"
	"fmt"
	"log"
	"os"
)

const AllPrivileges = 0777

var (
	inputFile, outputFile, keyFile, checkFile string
)

func CheckArgs() (err error) {
	if len(os.Args) != 5 {
		err = errors.New("bad args")
	} else {
		inputFile = os.Args[1]
		outputFile = os.Args[2]
		keyFile = os.Args[3]
		checkFile = os.Args[4]
	}
	return err
}

func main() {
	err := CheckArgs()
	if err == nil {
		data, _ := os.ReadFile(inputFile)

		key, _ := os.ReadFile(keyFile)
		key = desAlg.CompleteKey(key)

		encrypted := desAlg.Encrypt(data, desAlg.GenerateKeys(key))
		err = os.WriteFile(outputFile, encrypted, AllPrivileges)
		if err != nil {
			log.Fatalf("Can't write encrypted data, error is: %s", err)
		}

		decrypted := desAlg.Decrypt(encrypted, desAlg.GenerateKeys(key))
		_ = os.WriteFile(checkFile, decrypted[:len(data)], AllPrivileges)

		isEqual := bytes.Equal(data, decrypted[:len(data)])
		if isEqual {
			fmt.Printf("Input and decoded files are equal: %t\n", isEqual)
		} else {
			fmt.Printf("Bad result: %t\n", isEqual)
		}
	}
}
