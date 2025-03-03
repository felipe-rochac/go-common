package crypto

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha3"
	"crypto/sha512"
	"encoding/hex"
)

// MD5 returns the MD5 hash of the input string.
func MD5(input string) string {
	hash := md5.New()
	hash.Write([]byte(input))
	return hex.EncodeToString(hash.Sum(nil))
}

// HmacSha256 returns the HMAC-SHA-256 hash of the input string using the provided key.
func HmacSha256(key, input string) string {
	h := hmac.New(sha256.New, []byte(key))
	h.Write([]byte(input))
	return hex.EncodeToString(h.Sum(nil))
}

// HmacSha1 returns the HMAC-SHA-1 hash of the input string using the provided key.
func HmacSha1(key, input string) string {
	h := hmac.New(sha1.New, []byte(key))
	h.Write([]byte(input))
	return hex.EncodeToString(h.Sum(nil))
}

// HmacSha512 returns the HMAC-SHA-512 hash of the input string using the provided key.
func HmacSha512(key, input string) string {
	h := hmac.New(sha512.New, []byte(key))
	h.Write([]byte(input))
	return hex.EncodeToString(h.Sum(nil))
}

// Sha1 returns the Sha-1 hash of the input string.
func Sha1(input string) string {
	hash := sha1.New()
	hash.Write([]byte(input))
	return hex.EncodeToString(hash.Sum(nil))
}

// Sha256 returns the Sha-256 hash of the input string.
func Sha256(input string) string {
	hash := sha256.New()
	hash.Write([]byte(input))
	return hex.EncodeToString(hash.Sum(nil))
}

// Sha512 returns the Sha-512 hash of the input string.
func Sha512(input string) string {
	hash := sha512.New()
	hash.Write([]byte(input))
	return hex.EncodeToString(hash.Sum(nil))
}

// Sha3 returns the Sha3-256 hash of the input string.
func Sha3(input string) string {
	hash := sha3.New256()
	hash.Write([]byte(input))
	return hex.EncodeToString(hash.Sum(nil))
}
