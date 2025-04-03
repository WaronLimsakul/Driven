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

// func printRefreshToken(token database.RefreshToken) {
// 	fmt.Printf("Token: %v\n", token.Token)
// 	fmt.Printf("Revoked at: %v\n", token.RevokedAt)
// 	fmt.Printf("Revoked is not NULL: %v\n", token.RevokedAt.Valid)
// }

func ValidateRefreshToken(token database.RefreshToken) error {
	if time.Now().After(token.ExpiredAt) {
		return fmt.Errorf("refresh token expired")
	}

	// if valid = not null = revoked
	if token.RevokedAt.Valid {
		return fmt.Errorf("refresh token revoked")
	}
	return nil
}
