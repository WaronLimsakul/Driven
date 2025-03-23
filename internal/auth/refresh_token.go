package auth

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"time"

	"github.com/WaronLimsakul/Driven/internal/database"
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

func ValidateRefreshToken(token database.RefreshToken) error {
	if time.Now().After(token.ExpiredAt) {
		return fmt.Errorf("Refresh token expired")
	}

	if !token.RevokedAt.Valid {
		return fmt.Errorf("Refresh token revoked")
	}
	return nil
}
