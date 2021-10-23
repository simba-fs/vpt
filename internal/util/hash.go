package util

import (
	"crypto/sha256"
	"encoding/hex"
)

// Hash hashs input and return in hex form
func Hash(input string) string {
	a := sha256.Sum256([]byte(input))
	return hex.EncodeToString(a[:])
}
