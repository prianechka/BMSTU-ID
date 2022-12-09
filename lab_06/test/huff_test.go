package test

import (
	"bytes"
	huffman "lab_06/pkg"
	"log"
	"os"
	"testing"
)

func TestHuffman(t *testing.T) {
	files := []string{"test.txt", "im.png", "archive.zip", "main.out", "test.o", "image.jpg"}

	for _, file := range files {
		data, _ := os.ReadFile(file)
		tree := huffman.NewHuffmanTree(data)
		var compressedData []byte
		_ = tree.Compress(&compressedData)
		decompressedData, _ := tree.Decompress(compressedData)
		isEqual := bytes.Equal(data, decompressedData)
		if !isEqual {
			t.Errorf("not equal %s", file)
			log.Println(data[:10], decompressedData[:10])
		} else {
			log.Printf("file %s - passed", file)
		}
	}
}
