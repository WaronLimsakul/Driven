package auth

import (
	"testing"
	"time"

	"github.com/google/uuid"
)

// good unit test follows AAA patter (arrange-act-assert)
func TestCreateJWT(t *testing.T) {
	// arrange
	jwtSecret := "secret"

	type createJwtTestParams struct {
		name      string
		userID    uuid.UUID
		expiredIn time.Duration
		secret    string
		wantErr   bool
	}

	testCases := []createJwtTestParams{
		{"normal", uuid.New(), 3 * time.Minute, jwtSecret, false},
		{"past expired", uuid.New(), -3 * time.Minute, jwtSecret, false}, // want this for another validate testing
		{"no uuid", uuid.Nil, 3 * time.Minute, jwtSecret, true},
		{"no secret", uuid.New(), 3 * time.Minute, "", true},
	}

	// act
	for _, test := range testCases {
		_, err := CreateJWT(test.userID, test.expiredIn, test.secret)
		if test.wantErr && err == nil {
			t.Fatalf("test %s: expect error", test.name)
		} else if !test.wantErr && err != nil {
			t.Fatalf("test %s: expect: no error but get %v", test.name, err)
		}
	}
}

func TestValidateJWT(t *testing.T) {
	secret := "secret"

	type testValidateJwtParams struct {
		name          string
		userID        uuid.UUID
		expiredIn     time.Duration
		token         string
		shouldExpired bool
	}

	testCases := []testValidateJwtParams{
		{"normal", uuid.New(), 5 * time.Minute, "", false},
		{"normal expired token", uuid.New(), -5 * time.Minute, "", true},
	}

	for i, test := range testCases {
		token, err := CreateJWT(test.userID, test.expiredIn, secret)
		if err != nil {
			t.Fatalf("shouldn't happen at testValidateJWT")
		}
		testCases[i].token = token // test looks like is't
	}

	for _, test := range testCases {
		userID, err, expired := ValidateJWT(test.token, secret)
		if test.shouldExpired && (err == nil || !expired) {
			t.Fatalf("test %s: expect error and expired true found %v and %v", test.name, err, expired)
		} else if err == nil && test.userID != userID {
			t.Fatalf("test %s: expect user id %v but found %v", test.name, test.userID, userID)
		}
	}

}
