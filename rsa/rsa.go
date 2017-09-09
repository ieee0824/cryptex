package rsa

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha512"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
)

var (
	BadPrivateKeyErr = errors.New("bad private key")
	InvalidPubKeyErr = errors.New("invalid public key data")
	BadPublicKeyErr  = errors.New("not RSA public key")
)

func decodePrivateKey(key []byte) (*rsa.PrivateKey, error) {
	privateKeyBlock, _ := pem.Decode(key)
	if privateKeyBlock == nil {
		return nil, BadPrivateKeyErr
	}
	if privateKeyBlock.Type != "RSA PRIVATE KEY" {
		return nil, BadPrivateKeyErr
	}

	privateKey, err := x509.ParsePKCS1PrivateKey(privateKeyBlock.Bytes)
	if err != nil {
		return nil, err
	}

	return privateKey, err
}

func decodePublicKey(key []byte) (*rsa.PublicKey, error) {
	publicKeyBlock, _ := pem.Decode(key)
	if publicKeyBlock == nil {
		return nil, InvalidPubKeyErr
	}
	if publicKeyBlock.Type != "PUBLIC KEY" {
		return nil, errors.New(fmt.Sprintf("invalid public key type : %s", publicKeyBlock.Type))
	}

	publicKeyInterface, err := x509.ParsePKIXPublicKey(publicKeyBlock.Bytes)
	if err != nil {
		return nil, err
	}

	publicKey, ok := publicKeyInterface.(*rsa.PublicKey)
	if !ok {
		return nil, BadPublicKeyErr
	}

	return publicKey, nil
}

type RSA struct {
	privateKey []byte
	publicKey  []byte
	label      string
}

func New(pri, pub []byte) *RSA {
	return &RSA{
		privateKey: pri,
		publicKey:  pub,
	}
}

func (r *RSA) SetLabel(l string) *RSA {
	r.label = l
	return r
}

func (r *RSA) Encrypt(p []byte) ([]byte, error) {
	pubkey, err := decodePublicKey(r.publicKey)
	if err != nil {
		return nil, err
	}

	out, err := rsa.EncryptOAEP(sha512.New(), rand.Reader, pubkey, p, []byte(r.label))
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (r *RSA) Decrypt(c []byte) ([]byte, error) {
	priKey, err := decodePrivateKey(r.privateKey)
	if err != nil {
		return nil, err
	}

	out, err := rsa.DecryptOAEP(sha512.New(), rand.Reader, priKey, c, []byte(r.label))
	if err != nil {
		return nil, err
	}
	return out, nil
}
