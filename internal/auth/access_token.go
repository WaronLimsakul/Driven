package auth

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

// jwt is a data. If in string, it looks like this <method>.<payload>.<signature>
// the method and payload is only base 64 strings
// the signature is method(secret, message).
func CreateJWT(userID uuid.UUID, expiredIn time.Duration, secret string) (string, error) {
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

	// now we signed the <signature> part in token.
	signedToken, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

// just calculate method(secret, message) == expected_signature
func ValidateJWT(tokenString, secret string) (uuid.UUID, error) {
	// take the string, do the check in the function
	token, err := jwt.ParseWithClaims(
		tokenString,
		jwt.RegisteredClaims{},
		func(token *jwt.Token) (any, error) {
			// check if the method what I expect
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Token invalid: Unexpected sigin method: %v", token.Method.Alg())
			}

			return []byte(secret), nil
		})

	if err != nil {
		return uuid.UUID{}, fmt.Errorf("Validate token: Couldn't parse token: %v", err)
	}

	// check issuer
	issuer, err := token.Claims.GetIssuer()
	if err != nil {
		return uuid.UUID{}, fmt.Errorf("Validate token: No issuer found")
	} else if issuer != "Driven" {
		return uuid.UUID{}, fmt.Errorf("Validate token: Unexpected issuer: %v", issuer)
	}

	// check expire date
	expiredTime, err := token.Claims.GetExpirationTime()
	if err != nil {
		return uuid.UUID{}, fmt.Errorf("Validate token: Couldn't get expire time: %v", err)
	} else if time.Now().After(expiredTime.Time) {
		return uuid.UUID{}, fmt.Errorf("Validate token: the token is expired")
	}

	userID, err := token.Claims.GetSubject()
	if err != nil {
		return uuid.UUID{}, fmt.Errorf("Validate token: Couldn't get user id: %v", err)
	}

	userUUID, err := uuid.Parse(userID)
	if err != nil {
		return uuid.UUID{}, fmt.Errorf("Validate token: Couldn't parse user id: %v", err)
	}

	return userUUID, nil

}
