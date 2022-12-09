package digsig

import (
	"crypto/ed25519"
	"crypto/rand"
	"crypto/x509"
	"encoding/pem"
	"os"
)

func GenerateKeys(fileForPrivateKey, fileForPublicKey string) error {
	publicKey, privateKey, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		return err
	}

	privateKeyFile, createFileErr := os.Create(fileForPrivateKey)
	if createFileErr != nil {
		return err
	}
	defer privateKeyFile.Close()

	privateKeyBytes, marshalErr := x509.MarshalPKCS8PrivateKey(privateKey)
	if marshalErr != nil {
		return err
	}

	if encodingInFileErr := pem.Encode(privateKeyFile, &pem.Block{Type: "PRIVATE KEY", Bytes: privateKeyBytes}); encodingInFileErr != nil {
		return encodingInFileErr
	}

	publicKeyFile, createPublicFileErr := os.Create(fileForPublicKey)
	if createPublicFileErr != nil {
		return createPublicFileErr
	}
	defer publicKeyFile.Close()

	publicKeyBytes, marshalPublicKeyErr := x509.MarshalPKIXPublicKey(publicKey)
	if marshalPublicKeyErr != nil {
		return marshalPublicKeyErr
	}

	if encodePublicKeyErr := pem.Encode(publicKeyFile, &pem.Block{Type: "PUBLIC KEY", Bytes: publicKeyBytes}); encodePublicKeyErr != nil {
		return encodePublicKeyErr
	}

	return nil
}

func Sign(fileForPrivateKey, fileToSign, fileWithSign string) error {
	data, err := os.ReadFile(fileToSign)
	if err != nil {
		return err
	}

	key, readKeyFileErr := os.ReadFile(fileForPrivateKey)
	if readKeyFileErr != nil {
		return readKeyFileErr
	}

	decodedKey, _ := pem.Decode(key)
	privateKeyParsed, parseKeyErr := x509.ParsePKCS8PrivateKey(decodedKey.Bytes)
	if parseKeyErr != nil {
		return parseKeyErr
	}

	return os.WriteFile(fileWithSign, ed25519.Sign(privateKeyParsed.(ed25519.PrivateKey), data), 0644)
}

func Verify(fileWithPublicKey, fileToVerify, fileWithSignature string) (bool, error) {
	data, err := os.ReadFile(fileToVerify)
	if err != nil {
		return false, err
	}

	key, readPublicKeyErr := os.ReadFile(fileWithPublicKey)
	if readPublicKeyErr != nil {
		return false, readPublicKeyErr
	}

	signature, readSignatureErr := os.ReadFile(fileWithSignature)
	if readSignatureErr != nil {
		return false, readSignatureErr
	}

	decodedKey, _ := pem.Decode(key)
	pubKeyParsed, parseKeyError := x509.ParsePKIXPublicKey(decodedKey.Bytes)
	if parseKeyError != nil {
		return false, parseKeyError
	}

	return ed25519.Verify(pubKeyParsed.(ed25519.PublicKey), data, signature), nil
}
