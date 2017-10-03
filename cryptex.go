package cryptex

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"github.com/ghodss/yaml"
	"github.com/ieee0824/cryptex/encryptor"
	"io/ioutil"
	"os"
	"os/exec"
)

const (
	JSON = "json"
	YAML = "yaml"
)

var (
	NonePager       = errors.New("none pager")
	NotFoundCommand = errors.New("not found command")
)

type Container struct {
	EncryptionType string      `json:"encryption_type"`
	Values         interface{} `json:"values"`
}

func getEditor() string {
	if e := os.Getenv("DEFAULT_EDITOR"); e != "" {
		return e
	}

	if e := os.Getenv("EDITOR"); e != "" {
		return e
	}

	return "nano"
}

func hasCommand(c string) error {
	cmd := exec.Command("type", c)
	err := cmd.Run()
	if err != nil {
		return NotFoundCommand
	}
	return nil
}

func getPager() (string, error) {
	if p := os.Getenv("DEFAULT_PAGER"); p != "" {
		if err := hasCommand(p); err != nil {
			return "", err
		}
		return p, nil
	}

	if p := os.Getenv("PAGER"); p != "" {
		if err := hasCommand(p); err != nil {
			return "", err
		}
		return p, nil
	}

	if hasCommand("less") == nil {
		return "less", nil
	} else if hasCommand("more") == nil {
		return "more", nil
	} else if hasCommand("cat") == nil {
		return "cat", nil
	}

	return "", NonePager
}

type V struct {
	Value interface{}
}

type Cryptex struct {
	e          encryptor.Encryptor
	viewFormat string
}

func New(e encryptor.Encryptor) *Cryptex {
	c := &Cryptex{
		e:          e,
		viewFormat: "json",
	}
	return c
}

func (c *Cryptex) SetFormat(format string) *Cryptex {
	c.viewFormat = format
	return c
}

func (c *Cryptex) encrypt(obj interface{}) (interface{}, error) {
	m, ok := obj.(map[string]interface{})
	if !ok {
		v := &V{
			Value: obj,
		}
		bin, err := yaml.Marshal(v)
		if err != nil {
			return nil, err
		}

		result, err := c.e.Encrypt(bin)
		if err != nil {
			return nil, err
		}

		encoded := base64.StdEncoding.EncodeToString(result)

		return encoded, nil

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
		decoded, err := base64.StdEncoding.DecodeString(obj.(string))
		if err != nil {
			return nil, err
		}

		p, err := c.e.Decrypt(decoded)
		if err != nil {
			return nil, err
		}

		v := V{}

		if err := yaml.Unmarshal(p, &v); err != nil {
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

func (c *Cryptex) Encrypt(i interface{}) (*Container, error) {
	ret := &Container{
		EncryptionType: c.e.EncryptionType(),
	}
	o, err := c.encrypt(i)
	if err != nil {
		return nil, err
	}
	ret.Values = o
	return ret, nil
}

func (c *Cryptex) Decrypt(d *Container) (interface{}, error) {
	o, err := c.decrypt(d.Values)
	if err != nil {
		return nil, err
	}

	return o, nil
}

func (c *Cryptex) Edit(i *Container) (interface{}, error) {
	p, err := c.decrypt(i.Values)
	if err != nil {
		return nil, err
	}

	var bin []byte
	if c.viewFormat == "yaml" {
		var err error
		bin, err = yaml.Marshal(p)
		if err != nil {
			return nil, err
		}
	} else {
		var err error
		bin, err = json.MarshalIndent(p, "", "    ")
		if err != nil {
			return nil, err
		}
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

	if err := yaml.Unmarshal(pbin, &m); err != nil {
		return nil, err
	}

	result, err := c.encrypt(m)
	if err != nil {
		return nil, err
	}

	i.EncryptionType = c.e.EncryptionType()
	i.Values = result

	return i, nil
}

func (c *Cryptex) View(i *Container) error {
	pager, err := getPager()
	if err != nil {
		return err
	}
	p, err := c.decrypt(i.Values)
	if err != nil {
		return err
	}

	var bin []byte
	if c.viewFormat == "yaml" {
		var err error
		bin, err = yaml.Marshal(p)
		if err != nil {
			return err
		}
	} else {
		var err error
		bin, err = json.MarshalIndent(p, "", "    ")
		if err != nil {
			return err
		}
	}

	cmd := exec.Command(pager)
	cmd.Stdin = bytes.NewReader(bin)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}
