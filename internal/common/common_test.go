package common

import (
	"fmt"
	"testing"
)

func TestHashFunction(t *testing.T) {
	var str string

	str = "str"
	hash := HashFunction(str)
	fmt.Println(hash)

	str = "ctr"
	hash = HashFunction(str)
	fmt.Println(hash)
}
