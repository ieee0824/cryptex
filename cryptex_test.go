package cryptex

import (
	"reflect"
	"testing"
	"github.com/ghodss/yaml"
	"encoding/base64"
)

func genEncryptedValue(i interface{}) string {
	bin, _ := yaml.Marshal(V{i})
	for i, b := range bin {
		bin[i] = b + 1
	}



	return base64.StdEncoding.EncodeToString(bin)
}

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

func TestNew(t *testing.T) {
	c := New(&testEncrypter{})

	if c == nil {
		t.Fatalf("cannot allocate cryptex")
	}
}

func TestCryptex_Encrypt(t *testing.T) {
	tests := []struct {
		input map[string]interface{}
		want  map[string]interface{}
		err   bool
	}{
		{
			input: map[string]interface{}{},
			want:  map[string]interface{}{},
			err:   false,
		},
		{
			input: map[string]interface{}{
				"hoge": 1,
			},
			want: map[string]interface{}{
				"hoge": genEncryptedValue(1),
			},
			err: false,
		},
		{
			input: map[string]interface{}{
				"hoge": 1,
				"foo":  "bar",
				"map":  map[string]interface{}{},
			},
			want: map[string]interface{}{
				"hoge": genEncryptedValue(1),
				"foo":  genEncryptedValue("bar"),
				"map":  map[string]interface{}{},
			},
			err: false,
		},
		{
			input: map[string]interface{}{
				"hoge": 1,
				"foo":  "bar",
				"map": map[string]interface{}{
					"john":  "doe",
					"float": 1.1,
					"map": map[string]interface{}{
						"hoge": "huga",
					},
				},
			},
			want: map[string]interface{}{
				"hoge": genEncryptedValue(1),
				"foo":  genEncryptedValue("bar"),
				"map": map[string]interface{}{
					"john":  genEncryptedValue("doe"),
					"float": genEncryptedValue(1.1),
					"map": map[string]interface{}{
						"hoge": genEncryptedValue("huga"),
					},
				},
			},
			err: false,
		},
	}

	for _, test := range tests {
		got, err := New(&testEncrypter{}).Encrypt(test.input)
		if !test.err && err != nil {
			t.Fatalf("should not be error for %v but: %v", test.input, err)
		}
		if test.err && err == nil {
			t.Fatalf("should be error for %v but not:", test.input)
		}
		if !reflect.DeepEqual(got, test.want) {
			t.Fatalf("want %q, but %q:", test.want, got)
		}
	}
}

func TestCryptex_Decrypt(t *testing.T) {
	tests := []struct {
		input map[string]interface{}
		want  map[string]interface{}
		err   bool
	}{
		{
			input: map[string]interface{}{},
			want:  map[string]interface{}{},
		},
		{
			input: map[string]interface{}{
				"hoge": genEncryptedValue("huga"),
				"foo":  genEncryptedValue(1.1),
				"map": map[string]interface{}{
					"map": map[string]interface{}{
						"float": genEncryptedValue(1.1),
					},
					"john": genEncryptedValue("doe"),
				},
			},
			want: map[string]interface{}{
				"hoge": "huga",
				"foo":  1.1,
				"map": map[string]interface{}{
					"map": map[string]interface{}{
						"float": 1.1,
					},
					"john": "doe",
				},
			},
		},
	}

	for _, test := range tests {
		got, err := New(&testEncrypter{}).Decrypt(test.input)
		if !test.err && err != nil {
			t.Fatalf("should not be error for %v but: %v", test.input, err)
		}
		if test.err && err == nil {
			t.Fatalf("should be error for %v but not:", test.input)
		}
		if !reflect.DeepEqual(got, test.want) {
			t.Fatalf("want %q, but %q:", test.want, got)
		}
	}
}
