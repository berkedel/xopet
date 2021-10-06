package main

import (
	"encoding/binary"
	"fmt"
	"github.com/klauspost/compress/zstd"
	"io/ioutil"
	"log"
	"os"
)

var decoder, _ = zstd.NewReader(nil)

func List(zstdFilePath string) {
	plain, err := decompress(zstdFilePath)
	if err != nil {
		log.Fatalf("Error cannot decompress zst file: %s", err)
	}

	pivot := 0xe
	fileCount := int(binary.BigEndian.Uint32(plain[pivot : pivot+4]))
	pivot += 4

	for i := 1; i <= fileCount; i++ {
		fileNameSize := int(binary.BigEndian.Uint32(plain[pivot : pivot+4]))
		pivot += 4
		fileName := fmt.Sprintf(".%s", string(plain[pivot:pivot+fileNameSize]))
		pivot += fileNameSize
		dataAddress := binary.BigEndian.Uint32(plain[pivot : pivot+4])
		pivot += 4
		dataSize := binary.BigEndian.Uint32(plain[pivot : pivot+4])
		pivot += 4
		fmt.Printf("[%d] (0x%x - 0x%x) => %s\n", i, dataAddress, dataAddress+dataSize-1, fileName)
	}
}

func Unpack(zstdFilePath string, outDir string) {
	plain, err := decompress(zstdFilePath)
	if err != nil {
		log.Fatalf("Error cannot decompress zst file: %s", err)
	}

	pivot := 0xe
	fileCount := int(binary.BigEndian.Uint32(plain[pivot : pivot+4]))
	pivot += 4

	for i := 1; i <= fileCount; i++ {
		fileNameSize := int(binary.BigEndian.Uint32(plain[pivot : pivot+4]))
		pivot += 4
		fileName := string(plain[pivot : pivot+fileNameSize])
		pivot += fileNameSize
		dataAddress := binary.BigEndian.Uint32(plain[pivot : pivot+4])
		pivot += 4
		dataSize := binary.BigEndian.Uint32(plain[pivot : pivot+4])
		pivot += 4
		fmt.Printf("[%d] (0x%x - 0x%x) : .%s\n", i, dataAddress, dataAddress+dataSize-1, fileName)

		filePathName := fileName
		if len(outDir) > 0 {
			filePathName = fmt.Sprintf("%s%s", outDir, fileName)
		} else {
			filePathName = fmt.Sprintf("%s%s", ".", fileName)
		}
		createDir(filePathName)
		dumpFile(filePathName, plain[dataAddress:dataAddress+dataSize])
	}
}

func decompress(zstdFilePath string) ([]byte, error) {
	file, err := os.Open(zstdFilePath)
	if err != nil {
		log.Fatalf("Error reading zst file: %s.", err)
	}

	defer file.Close()

	content, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatalf("Error cannot load zst file: %s", err)
	}

	return decoder.DecodeAll(content, nil)
}
