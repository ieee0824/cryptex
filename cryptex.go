package cryptex

import (
	"encoding/json"
	"github.com/ieee0824/cryptex/encryptor"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

func getEditor() string {
	if e := os.Getenv("PNZR_EDITOR"); e != "" {
		return e
	}

	if e := os.Getenv("EDITOR"); e != "" {
		return e
	}

	return "nano"
}

type V struct {
	Value interface{}
}

type Cryptex struct {
	e   encryptor.Encryptor
}

func New(e encryptor.Encryptor) *Cryptex {
	c := &Cryptex{
		e:   e,
	}
	return c
}

func (c *Cryptex) encrypt(obj interface{}) (interface{}, error) {
	m, ok := obj.(map[string]interface{})
	if !ok {
		v := &V{
			Value: obj,
		}
		bin, err := json.Marshal(v)
		if err != nil {
			return nil, err
		}

		result, err := c.e.Encrypt(bin)
		if err != nil {
			return nil, err
		}

		return result, nil

	}

	for k, v := range m {
		o, err := c.encrypt(v)
		if err != nil {
			return nil, err
		}
		obj.(map[string]interface{})[k] = o
	}
	return obj, nil
}

func (c *Cryptex) decrypt(obj interface{}) (interface{}, error) {
	m, ok := obj.(map[string]interface{})
	if !ok {
		p, err := c.e.Decrypt(obj.([]byte))
		if err != nil {
			log.Fatalln(err)
			return nil, err
		}

		v := V{}

		if err := json.Unmarshal(p, &v); err != nil {
			return nil, err
		}

		return v.Value, nil
	}

	for k, v := range m {
		o, err := c.decrypt(v)
		if err != nil {
			return nil, err
		}
		obj.(map[string]interface{})[k] = o
	}
	return obj, nil
}

func (c *Cryptex) Encrypt(i interface{}) (interface{}, error) {
	o, err := c.encrypt(i)
	if err != nil {
		return nil, err
	}
	return o, nil
}

func (c *Cryptex) Decrypt(d interface{}) (interface{}, error) {
	o, err := c.decrypt(d)
	if err != nil {
		return nil, err
	}

	return o, nil
}

func (c *Cryptex) Edit(i interface{}) (interface{}, error) {
	p, err := c.decrypt(i)
	if err != nil {
		return nil, err
	}
	bin, err := json.MarshalIndent(p, "", "    ")
	if err != nil {
		return nil, err
	}

	file, err := ioutil.TempFile(os.TempDir(), "cryptex")
	if err != nil {
		return nil, err
	}
	defer os.Remove(file.Name())
	if _, err := file.Write(bin); err != nil {
		return nil, err
	}

	cmd := exec.Command(getEditor(), file.Name())
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return nil, err
	}

	pbin, err := ioutil.ReadFile(file.Name())
	if err != nil {
		return nil, err
	}

	m := map[string]interface{}{}

	if err := json.Unmarshal(pbin, &m); err != nil {
		return nil, err
	}

	result, err := c.encrypt(m)
	if err != nil {
		return nil, err
	}

	return result, nil
}
