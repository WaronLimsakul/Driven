package auth

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

// jwt is a data. If in string, it looks like this <method>.<payload>.<signature>
// the method and payload is only base 64 strings
// the signature is method(secret, message).

// this function create jwt with 4 claims: issuer, issued at, expired at, and user id
func CreateJWT(userID uuid.UUID, expiredIn time.Duration, secret string) (string, error) {
	if userID == uuid.Nil {
		return "", fmt.Errorf("user ID not found")
	}
	// claims are the <payload> part
	// RegisteredClaims is standard claim lib gives us. Great for access token.
	claims := jwt.RegisteredClaims{
		Issuer:    "Driven",
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(expiredIn)),
		Subject:   userID.String(),
	}

	// create a token with method and claims
	// We choose HS256 = HMAC + SHA-256 = use shared secret key
	// Good for monolith server (if same service sign and verify it).
	// RSA and ECDSA are good for microservices, but we are not them.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	if secret == "" {
		return "", fmt.Errorf("secret is empty")
	}
	// now we signed the <signature> part in token.
	signedToken, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

// Just calculate method(secret, message) == expected_signature.
// Return the user uuid, err, and isExpired (bool) if the token is valid
func ValidateJWT(tokenString, secret string) (uuid.UUID, error, bool) {
	// take the string, do the check in the function
	token, err := jwt.ParseWithClaims(
		tokenString,
		&jwt.RegisteredClaims{},
		func(token *jwt.Token) (any, error) {
			// check if the method what I expect
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Token invalid: Unexpected sigin method: %v", token.Method.Alg())
			}

			return []byte(secret), nil
		})

	isExpired := false
	// this parse token already check the expired time
	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			isExpired = true
		}
		return uuid.UUID{}, fmt.Errorf("Validate token: Couldn't parse token: %v", err), isExpired
	}

	// check issuer
	issuer, err := token.Claims.GetIssuer()
	if err != nil {
		return uuid.UUID{}, fmt.Errorf("Validate token: No issuer found"), false
	} else if issuer != "Driven" {
		return uuid.UUID{}, fmt.Errorf("Validate token: Unexpected issuer: %v", issuer), false
	}

	userID, err := token.Claims.GetSubject()
	if err != nil {
		return uuid.UUID{}, fmt.Errorf("Validate token: Couldn't get user id: %v", err), false
	}

	userUUID, err := uuid.Parse(userID)
	if err != nil {
		return uuid.UUID{}, fmt.Errorf("Validate token: Couldn't parse user id: %v", err), false
	}

	return userUUID, nil, false
}
