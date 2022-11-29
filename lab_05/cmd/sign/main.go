package main

import (
	"flag"
	digsig "lab_05/pkg"
	"log"
)

var (
	fileForPrivateKey, fileToSign, fileWithSign string
)

func init() {
	flag.StringVar(&fileForPrivateKey, "priv", "privkey.pem", "private key filename")
	flag.StringVar(&fileToSign, "file", "test.txt", "file to sign")
	flag.StringVar(&fileWithSign, "sig", "signature.sig", "file to save signature")
}

func main() {
	flag.Parse()
	if err := digsig.Sign(fileForPrivateKey, fileToSign, fileWithSign); err != nil {
		log.Fatalf("Failed to sign file, error is: %s", err)
	}
}
