package pkg

import (
	"crypto/sha256"
	"fmt"
	"io"
	"log"
)

func HashWithReader(reader io.Reader) (string, error) {
	data, err := io.ReadAll(reader)
	if err != nil {
		log.Printf("hashing error %v", err)
		return "", err
	}
	hash := sha256.Sum256(data)
	return fmt.Sprintf("%x", hash), nil
}

func HashWithString(data string) string {
	hash := sha256.Sum256([]byte(data))
	return fmt.Sprintf("%x", hash)
}
