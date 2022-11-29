package test

import (
	"bytes"
	desAlg "desAlg/pkg"
	"log"
	"os"
	"testing"
)

func TestDesFunction(t *testing.T) {
	files := []string{"im.png", "archive.zip", "main.out", "test.o"}

	for _, file := range files {
		data, _ := os.ReadFile(file)
		key, _ := os.ReadFile("key.txt")
		key = desAlg.CompleteKey(key)

		encrypted := desAlg.Encrypt(data, desAlg.GenerateKeys(key))
		decrypted := desAlg.Decrypt(encrypted, desAlg.GenerateKeys(key))

		isEqual := bytes.Equal(data, decrypted[:len(data)])
		if !isEqual {
			t.Errorf("not equal %s", file)
		} else {
			log.Printf("file %s - passed", file)
		}
	}
}
