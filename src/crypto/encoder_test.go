package crypto

import (
	"encoding/base64"
	"testing"
)

const helloWorld string = "Hello, World!"

func TestBase64Encode(t *testing.T) {
	input := []byte(helloWorld)
	want := base64.StdEncoding.EncodeToString(input)
	got := Base64Encode(input)
	if got != want {
		t.Errorf("Base64Encode() = %v, want %v", got, want)
	}
}

func TestBase64Decode(t *testing.T) {
	input := base64.StdEncoding.EncodeToString([]byte(helloWorld))
	want := []byte(helloWorld)
	got, err := Base64Decode(input)
	if err != nil {
		t.Errorf("Base64Decode() error = %v", err)
	}
	if string(got) != string(want) {
		t.Errorf("Base64Decode() = %v, want %v", got, want)
	}
}

func TestBase64Decode_InvalidInput(t *testing.T) {
	input := "Invalid base64 string"
	_, err := Base64Decode(input)
	if err == nil {
		t.Errorf("Base64Decode() expected error for invalid input, got nil")
	}
}
