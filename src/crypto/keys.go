package crypto

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/rsa"
	"encoding/pem"
	"os"
)

// GenerateKeyPair generates a new RSA 2048 public and private key pair.
func GenerateKeyPairRSA() (*rsa.PrivateKey, *rsa.PublicKey, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, nil, err
	}
	return privateKey, &privateKey.PublicKey, nil
}

// GenerateKeyPair generates a new ECDSA P256 public and private key pair.
func GenerateKeyPairECDSA() (*ecdsa.PrivateKey, *ecdsa.PublicKey, error) {
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return nil, nil, err
	}
	return privateKey, &privateKey.PublicKey, nil
}

// SavePEMKey saves a private key to a PEM file.
func SavePEMKey(fileName string, key []byte) error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	var pemKey = &pem.Block{
		Type:  "EC PRIVATE KEY",
		Bytes: key,
	}

	return pem.Encode(file, pemKey)
}

// SavePublicPEMKey saves a public key to a PEM file.
func SavePublicPEMKey(fileName string, pubkey []byte) error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	var pemKey = &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: pubkey,
	}

	return pem.Encode(file, pemKey)
}
