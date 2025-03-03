package crypto

import "encoding/base64"

// Base64Encode encodes the input byte array to a base64 string.
func Base64Encode(input []byte) string {
	return base64.StdEncoding.EncodeToString(input)
}

// Base64Decode decodes the input base64 string to a byte array.
func Base64Decode(input string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(input)
}
