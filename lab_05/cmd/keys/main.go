package main

import (
	"flag"
	digsig "lab_05/pkg"
	"log"
)

var (
	fileForPrivateKey, fileForPublicKey string
)

func init() {
	flag.StringVar(&fileForPrivateKey, "priv", "privkey.pem", "private key filename")
	flag.StringVar(&fileForPublicKey, "pub", "pubkey.pem", "public key filename")
}

func main() {
	flag.Parse()
	if err := digsig.GenerateKeys(fileForPrivateKey, fileForPublicKey); err != nil {
		log.Fatalf("Failed to generate keys, error is: %s", err)
	}
}
