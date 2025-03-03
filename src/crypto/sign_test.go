package crypto

import (
	"testing"
)

func TestGenerateCronSignature(t *testing.T) {
	secret := "mysecret"
	messages := []string{"message1", "message2"}
	expectedSignature := "d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2"

	signature, err := GenerateCronSignature(secret, messages)
	if err != nil {
		t.Fatalf("GenerateCronSignature() error = %v", err)
	}
	if signature != expectedSignature {
		t.Errorf("GenerateCronSignature() = %v, want %v", signature, expectedSignature)
	}
}

func TestValidateCronSignature(t *testing.T) {
	secret := "mysecret"
	messages := []string{"message1", "message2"}
	signature := "d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2"

	valid, err := ValidateCronSignature(secret, signature, messages)
	if err != nil {
		t.Fatalf("ValidateCronSignature() error = %v", err)
	}
	if !valid {
		t.Errorf("ValidateCronSignature() = %v, want %v", valid, true)
	}

	invalidSignature := "invalidsignature"
	valid, err = ValidateCronSignature(secret, invalidSignature, messages)
	if err != nil {
		t.Fatalf("ValidateCronSignature() error = %v", err)
	}
	if valid {
		t.Errorf("ValidateCronSignature() = %v, want %v", valid, false)
	}
}
