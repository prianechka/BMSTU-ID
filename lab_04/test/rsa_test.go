package test

import (
	"bytes"
	alg "lab_04/pkg"
	"log"
	"os"
	"testing"
)

func TestRSA(t *testing.T) {
	var limit int64 = 2048
	RSA := alg.NewRSA(limit)
	files := []string{"im.png", "archive.zip", "main.out", "test.o", "image.jpg"}

	for _, file := range files {
		data, _ := os.ReadFile(file)
		encryptedData := RSA.Encrypt(data)
		result := RSA.Decrypt(encryptedData)
		isEqual := bytes.Equal(data, result)
		if !isEqual {
			t.Errorf("not equal %s", file)
			log.Println(data[:10], result[:10])
		} else {
			log.Printf("file %s - passed", file)
		}
	}
}
