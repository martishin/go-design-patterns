package main

import (
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/martishin/go-design-patterns/patterns/structural/decorator/internal/decorator"
	"github.com/martishin/go-design-patterns/patterns/structural/decorator/internal/filedata"
)

func main() {
	path := filepath.Join(os.TempDir(), "decorator-demo.dat")
	key := []byte("0123456789abcdef0123456789abcdef")
	payload := []byte("name,salary\nAlice,120000\nBob,95000")

	fileSource := filedata.NewFileDataSource(path)

	encryptedSource, err := decorator.NewEncryptionDecorator(fileSource, key)
	if err != nil {
		log.Fatal(err)
	}

	compressedAndEncryptedSource, err := decorator.NewCompressionDecorator(encryptedSource)
	if err != nil {
		log.Fatal(err)
	}

	if err := compressedAndEncryptedSource.Write(payload); err != nil {
		log.Fatal(err)
	}

	storedBytes, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	recoveredPayload, err := compressedAndEncryptedSource.Read()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Stored file: %s\n", path)
	fmt.Printf("Stored bytes (hex): %s\n", hex.EncodeToString(storedBytes))
	fmt.Printf("Recovered payload:\n%s\n", recoveredPayload)
}
