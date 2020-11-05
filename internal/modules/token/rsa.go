package token

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"github.com/dgrijalva/jwt-go"
	"strings"
)

type RSA struct {
	privateKey *rsa.PrivateKey
}

func newRSA(privateKey string) *RSA {
	if privateKey == "" {
		pk, _ := NewRSAPrivateKey(1024)
		privateKey = base64.StdEncoding.EncodeToString(pk.PEM())
	}

	pk, err := base64.StdEncoding.DecodeString(privateKey)
	if err != nil {
		panic(err)
	}

	jwtKey, err := jwt.ParseRSAPrivateKeyFromPEM(pk)
	if err != nil {
		panic(err)
	}

	return &RSA{
		privateKey: jwtKey,
	}
}

func (r *RSA) PrivateKey() *rsa.PrivateKey {
	return r.privateKey
}

type RSAPrivateKey struct {
	PrivateKey *rsa.PrivateKey
	Seed       []byte
}

func NewRSAPrivateKey(bits int) (*RSAPrivateKey, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return nil, err
	}
	return &RSAPrivateKey{PrivateKey: privateKey}, nil
}

func (pk *RSAPrivateKey) Encrypt(in []byte) ([]byte, error) {
	return rsa.EncryptOAEP(sha1.New(), bytes.NewReader(appendSeed(pk.Seed)), &pk.PrivateKey.PublicKey, in, nil)
}

func (pk *RSAPrivateKey) Decrypt(in []byte) ([]byte, error) {
	return rsa.DecryptOAEP(sha1.New(), nil, pk.PrivateKey, in, nil)
}

func (pk *RSAPrivateKey) PEM() []byte {
	derStream := x509.MarshalPKCS1PrivateKey(pk.PrivateKey)

	return pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: derStream,
	})
}

func (pk *RSAPrivateKey) PublicPEM() []byte {
	derPkix, _ := x509.MarshalPKIXPublicKey(&pk.PrivateKey.PublicKey)

	return pem.EncodeToMemory(&pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: derPkix,
	})
}

func appendSeed(seed []byte) []byte {
	n := len(seed)
	if n >= 20 {
		return seed[0:20]
	}
	return append(seed, []byte(strings.Repeat("0", 20-n))...)
}
