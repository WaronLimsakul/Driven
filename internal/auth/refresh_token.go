package auth

import (
	"crypto/rand"
	"encoding/hex"
)

func CreateRefreshToken() (string, error) {
	// 32 bytes: long enough to avoid collision + bruteforce, short enough
	// to store and carry around.
	bytesToken := make([]byte, 32) // zero init

	// put random things in those bytes
	_, err := rand.Read(bytesToken)
	if err != nil {
		return "", err
	}

	// store as hex because every computer have 0-9 a-f
	token := hex.EncodeToString(bytesToken)
	return token, nil
}
