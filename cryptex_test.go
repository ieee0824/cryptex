package cryptex

type testEncrypter struct {

}

func (t *testEncrypter) Encrypt(d []byte) ([]byte, error) {
	for i, v := range d {
		d[i] = v + 1
	}

	return d, nil
}

func (t *testEncrypter) Decrypt(d []byte) ([]byte, error) {
	for i, v := range d {
		d[i] = v - 1
	}

	return d, nil
}
