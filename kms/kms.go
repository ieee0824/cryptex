package kms

import (
	"encoding/json"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kms"
	"github.com/aws/aws-sdk-go/service/kms/kmsiface"
)

func split(b []byte, n int) [][]byte {
	size := len(b)
	if b == nil {
		return nil
	}
	if size == 0 {
		return [][]byte{[]byte{}}
	}
	if size <= n {
		return [][]byte{b}
	}
	var ret [][]byte
	if size%n == 0 {
		ret = make([][]byte, 0, size/n)
	} else {
		ret = make([][]byte, 0, size/n+1)
	}

	for i := 0; i < size; i += n {
		end := i + n
		if size < end {
			end = size
		}
		ret = append(ret, b[i:end])
	}

	return ret
}

type KMS struct {
	keyID *string
	svc   kmsiface.KMSAPI
}

func New(sess *session.Session) *KMS {
	return &KMS{svc: kms.New(sess)}
}

func (k *KMS) EncryptionType() string {
	return "kms"
}

func (k *KMS) SetKey(key string) *KMS {
	k.keyID = &key

	return k
}

func (k *KMS) Encrypt(p []byte) ([]byte, error) {
	var buffer [][]byte
	for _, p := range split(p, 2048) {
		params := &kms.EncryptInput{
			KeyId:     k.keyID,
			Plaintext: p,
		}

		resp, err := k.svc.Encrypt(params)
		if err != nil {
			return nil, err
		}

		buffer = append(buffer, resp.CiphertextBlob)
	}

	return json.Marshal(buffer)
}

func (k *KMS) Decrypt(c []byte) ([]byte, error) {
	var buffer [][]byte
	var ret []byte
	if err := json.Unmarshal(c, &buffer); err != nil {
		// Processing for backward compatibility.
		buffer = [][]byte{c}
	}

	for _, c := range buffer {
		params := &kms.DecryptInput{
			CiphertextBlob: c,
		}

		resp, err := k.svc.Decrypt(params)
		if err != nil {
			return nil, err
		}
		ret = append(ret, resp.Plaintext...)
	}

	return ret, nil
}
