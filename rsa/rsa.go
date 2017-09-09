package rsa

import (
	"encoding/pem"
	"fmt"
	"crypto/x509"
	"crypto/rsa"
	"crypto/sha512"
	"errors"
	"crypto/rand"
	"crypto/sha1"
)

var (
	NotPemEncodeErr = errors.New("not PEM-encoded")
	UnknownKeyTypeErr = errors.New("Unknown key type error")
	BadPrivateKeyErr = errors.New("bad private key")
	InvalidPubKeyErr = errors.New("invalid public key data")
	BadPublicKeyErr = errors.New("not RSA public key")
)

func decodePublicKey(key []byte) (*rsa.PublicKey, error) {
	publicKeyBlock, _ := pem.Decode(key)
	if publicKeyBlock == nil {
		return nil,InvalidPubKeyErr
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
	publicKey []byte
}

func New(pri, pub []byte) *RSA {
	return &RSA{
		privateKey: pri,
		publicKey: pub,
	}
}

func (r *RSA) Encrypt(p []byte) ([]byte, error) {
	pubkey, err := decodePublicKey(r.publicKey)
	if err != nil {
		return nil, err
	}

	out, err := rsa.EncryptOAEP(sha512.New(), rand.Reader, pubkey, p, []byte(""))
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (r *RSA) Decrypt(c []byte) ([]byte, error) {
block, _ := pem.Decode(r.privateKey)
	if block == nil {
		return nil, NotPemEncodeErr
	}
	got, want := block.Type, "RSA PRIVATE KEY"
	if  got != want {
		return nil, errors.New(fmt.Sprintf("%v: %q, want: %q", UnknownKeyTypeErr.Error(), got, want))
	}
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, BadPrivateKeyErr
	}
	out, err := rsa.DecryptOAEP(sha1.New(), rand.Reader, priv, c, []byte(""))
	if err != nil {
		return nil, err
	}
	return out, nil
}