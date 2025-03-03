package crypto

import (
	"crypto/hmac"

	"crypto/sha256"
	"encoding/hex"
)

// GenerateCronSignature generates a cron signature using HMAC with SHA256
func GenerateCronSignature(key string, parameters []string) (string, error) {
	h := hmac.New(sha256.New, []byte(key))
	for _, parameter := range parameters {
		h.Write([]byte(parameter))
	}
	return hex.EncodeToString(h.Sum(nil)), nil
}

// ValidateCronSignature validates a cron signature using HMAC with SHA256
func ValidateCronSignature(secret, signature string, parameters []string) (bool, error) {
	expectedSignature, err := GenerateCronSignature(secret, parameters)
	if err != nil {
		return false, err
	}
	return hmac.Equal([]byte(expectedSignature), []byte(signature)), nil
}
