package main

import (
	"flag"
	huffman "lab_06/pkg"
	"log"
	"os"
)

var (
	file, compressedFile, decompressFile, realCompressedFile string
)

func init() {
	flag.StringVar(&file, "f", "test.txt", "file to compress")
	flag.StringVar(&compressedFile, "c", "compressed.txt", "file to store compressed result")
	flag.StringVar(&decompressFile, "d", "decompressed.txt", "file to store decompressed result")
	flag.StringVar(&realCompressedFile, "k", "result.bin", "file with real size")
}

func main() {

	flag.Parse()

	srcFileData, err := os.ReadFile(file)
	if err != nil {
		log.Fatalf("Can't open file, error is: %s", err)
	}

	tree := huffman.NewHuffmanTree(srcFileData)
	var compressedData []byte

	if compressErr := tree.Compress(&compressedData); compressErr != nil {
		log.Fatalf("Failed to compress, error is: %s", compressErr)
	}
	if writeCompressErr := os.WriteFile(compressedFile, compressedData, 0644); writeCompressErr != nil {
		log.Fatalf("Can't write compressed data, error is: %s", writeCompressErr)
	}

	bytesArray := huffman.MakeArrayOfBytes(compressedData, 8)
	result := huffman.CreateBytesFromBits(bytesArray)
	if writeRealErr := os.WriteFile(realCompressedFile, result, 0644); writeRealErr != nil {
		log.Fatalf("Can't write data, error is: %s", writeRealErr)
	}

	decompressedData, decompressErr := tree.Decompress(compressedData)
	if decompressErr != nil {
		log.Fatalf("Failed to decompress, error is: %s", decompressErr)
	}
	if writeResultErr := os.WriteFile(decompressFile, decompressedData, 0644); writeResultErr != nil {
		log.Fatalf("Can't write decompressed data, error is: %s", err)
	}
}
