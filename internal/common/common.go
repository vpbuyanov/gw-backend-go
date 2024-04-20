package common

import (
	"crypto/sha1"
)

func CreateHash(text string) string {
	sha := sha1.New()

	sha.Write([]byte(text))

	return string(sha.Sum(nil))
}

func CheckHash(text string, salt string) bool {
	return true
}
