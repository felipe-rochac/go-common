package crypto

import (
	"crypto/x509"
	"os"
	"testing"
)

func TestGenerateKeyPairRSA(t *testing.T) {
	privateKey, publicKey, err := GenerateKeyPairRSA()
	if err != nil {
		t.Fatalf("GenerateKeyPairRSA() error = %v", err)
	}
	if privateKey == nil || publicKey == nil {
		t.Fatalf("GenerateKeyPairRSA() returned nil keys")
	}
}

func TestGenerateKeyPairECDSA(t *testing.T) {
	privateKey, publicKey, err := GenerateKeyPairECDSA()
	if err != nil {
		t.Fatalf("GenerateKeyPairECDSA() error = %v", err)
	}
	if privateKey == nil || publicKey == nil {
		t.Fatalf("GenerateKeyPairECDSA() returned nil keys")
	}
}

func TestSavePEMKey(t *testing.T) {
	privateKey, _, err := GenerateKeyPairECDSA()
	if err != nil {
		t.Fatalf("GenerateKeyPairECDSA() error = %v", err)
	}

	keyBytes, err := x509.MarshalECPrivateKey(privateKey)
	if err != nil {
		t.Fatalf("x509.MarshalECPrivateKey() error = %v", err)
	}

	fileName := "test_private_key.pem"
	defer os.Remove(fileName)

	err = SavePEMKey(fileName, keyBytes)
	if err != nil {
		t.Fatalf("SavePEMKey() error = %v", err)
	}

	_, err = os.Stat(fileName)
	if os.IsNotExist(err) {
		t.Fatalf("SavePEMKey() did not create file")
	}
}

func TestSavePublicPEMKey(t *testing.T) {
	_, publicKey, err := GenerateKeyPairECDSA()
	if err != nil {
		t.Fatalf("GenerateKeyPairECDSA() error = %v", err)
	}

	keyBytes, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		t.Fatalf("x509.MarshalPKIXPublicKey() error = %v", err)
	}

	fileName := "test_public_key.pem"
	defer os.Remove(fileName)

	err = SavePublicPEMKey(fileName, keyBytes)
	if err != nil {
		t.Fatalf("SavePublicPEMKey() error = %v", err)
	}

	_, err = os.Stat(fileName)
	if os.IsNotExist(err) {
		t.Fatalf("SavePublicPEMKey() did not create file")
	}
}
