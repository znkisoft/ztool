package util

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
)

const (
	secret = "07M5oF8hwcp4" // TEMP: refactor it when you see it.
)

func HashPassword(password string) (string, error) {
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(password))
	sha := hex.EncodeToString(mac.Sum(nil))
	return sha, nil
}

func CheckPasswordHash(password, hash string) bool {
	p, _ := HashPassword(password)
	return hash == p
}

/**
// ValidMAC reports whether messageMAC is a valid HMAC tag for message.
func ValidMAC(message, messageMAC, key []byte) bool {
	mac := hmac.New(sha256.New, key)
	mac.Write(message)
	expectedMAC := mac.Sum(nil)
	return hmac.Equal(messageMAC, expectedMAC)
}
*/
