package auth

import (
	"errors"
	"time"
)

type NagbotClaims struct {
	UserId  int   `json:"uid,omitempty"` // User ID for the authenticated user.
	SessId  int   `json:"sid,omitempty"` // Session ID, used to expire all outstanding sessions at once.
	Issued  int64 `json:"iss,omitempty"` // Unix time of server at time of issuing
	Expires int64 `json:"exp,omitempty"` // If 0, does not expire (or up to server to decide)
}

var (
	InvalidExpiryDateError = errors.New("Expiry date is before issued date.")
	InvalidIssuedDateError = errors.New("Issued date is in the future.")
	ExpiredTokenError      = errors.New("Token is expired.")
)

func (c NagbotClaims) Valid() error {

	// Is an expiry time set?
	if c.Expires != 0 {
		now := time.Now().Unix()

		// Check for invalid timestamps.
		if c.Expires < c.Issued {
			return InvalidExpiryDateError
		}

		if now < c.Issued {
			return InvalidIssuedDateError
		}

		// Check for expired tokens.
		if now >= c.Expires {
			return ExpiredTokenError
		}

	}

	return nil
}
