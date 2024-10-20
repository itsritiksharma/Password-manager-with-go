package hash

import (
	"crypto/sha256"
	"fmt"
)

func Hash(password []byte) string {
	h := sha256.New()

	h.Write([]byte(password))

	pass := h.Sum(nil)

	fmt.Println("h", h)

	fmt.Println("Pass", pass)

	return "asdf"
}
