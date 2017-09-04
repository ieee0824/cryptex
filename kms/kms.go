package kms

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kms"
	"github.com/aws/aws-sdk-go/service/kms/kmsiface"
)

type KMS struct {
	keyID *string
	svc   kmsiface.KMSAPI
}

func New(sess *session.Session) *KMS {
	return &KMS{svc: kms.New(sess)}
}

func (k *KMS) SetKey(key string) *KMS {
	k.keyID = &key

	return k
}

func (k *KMS) Encrypt(p []byte) ([]byte, error) {
	params := &kms.EncryptInput{
		KeyId:     k.keyID,
		Plaintext: p,
	}

	resp, err := k.svc.Encrypt(params)
	if err != nil {
		return nil, err
	}

	return resp.CiphertextBlob, nil
}

func (k *KMS) Decrypt(c []byte) ([]byte, error) {
	params := &kms.DecryptInput{
		CiphertextBlob: c,
	}

	resp, err := k.svc.Decrypt(params)
	if err != nil {
		return nil, err
	}

	return resp.Plaintext, nil
}
