package common

import (
	"crypto/sha1"
)

func HashFunction(text string) string {
	sha := sha1.New()

	sha.Write([]byte(text))

	return string(sha.Sum(nil))
}
