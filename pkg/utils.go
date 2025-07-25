package pkg

import (
	"crypto/sha256"
	"fmt"
	"io"
	"log"
	"net/http"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
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

func Capital(val any) string {
	str := fmt.Sprintf("%v", val)
	return cases.Title(language.English).String(str)
}

func IsHTMXRequest(r *http.Request) bool {
	return r.Header.Get("HX-Request") == "true"
}
