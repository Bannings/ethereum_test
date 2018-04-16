package sig

import (
	"crypto"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
)

func VerifySHA256RSA(pubPEM string, message []byte, sig string) error {
	pub, err := ParsePKIXPublicKey(pubPEM)
	if err != nil {
		return err
	}

	hashed := sha256.Sum256(message)
	signed, err := base64.StdEncoding.DecodeString(sig)
	if err != nil {
		return err
	}
	return rsa.VerifyPKCS1v15(pub, crypto.SHA256, hashed[:], signed)
}

func ParsePKIXPublicKey(pubPEM string) (*rsa.PublicKey, error) {
	block, _ := pem.Decode([]byte(pubPEM))
	if block == nil {
		return nil, errors.New("failed to parse PEM block containing the public key")
	}

	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	switch pub := pub.(type) {
	case *rsa.PublicKey:
		return pub, nil
	default:
		return nil, errors.New("unknown type of public key")
	}
}
