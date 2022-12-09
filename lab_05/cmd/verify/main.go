package main

import (
	"flag"
	"fmt"
	digsig "lab_05/pkg"
	"log"
)

var (
	fileWithPublicKey, fileToVerify, fileWithSignature string
)

func init() {
	flag.StringVar(&fileWithPublicKey, "pub", "pubkey.pem", "public key filename")
	flag.StringVar(&fileToVerify, "file", "test.txt", "file to verify")
	flag.StringVar(&fileWithSignature, "sig", "signature.sig", "signature file")
}

func main() {
	flag.Parse()
	isValid, err := digsig.Verify(fileWithPublicKey, fileToVerify, fileWithSignature)
	if err != nil {
		log.Fatalf("Failed to sign file, error is: %s", err)
	}
	if !isValid {
		fmt.Println("Signature is corrupted")
	} else {
		fmt.Println("Signature is correct")
	}
}
