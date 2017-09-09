package main

import (
	"github.com/ieee0824/cryptex/rsa"
	"io/ioutil"
	"log"
	"github.com/ieee0824/cryptex"
	"fmt"
	"encoding/json"
)

func main() {
	priBin, err := ioutil.ReadFile("private-key.pem")
	if err != nil {
		log.Fatalln(err)
	}
	pubBin, err := ioutil.ReadFile("public-key.pem")
	if err != nil {
		log.Fatalln(err)
	}

	encryptor := rsa.New(priBin, pubBin)
	plainMap := map[string]interface{}{
		"hoge":  "huga",
		"foo":   "bar",
		"int":   0,
		"float": 1.1,
		"sub_map": map[string]interface{}{
			"alice": 12,
			"bob":   25,
		},
	}
	c := cryptex.New(encryptor)
	if err != nil {
		log.Fatalln(err)
	}
	cipher, err := c.Encrypt(plainMap)
	if err != nil {
		log.Fatalln(err)
	}
	binEncrypted, err := json.MarshalIndent(cipher, "", "    ")
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(string(binEncrypted))

	plain, err := c.Decrypt(cipher)
	if err != nil {
		log.Fatalln(err)
	}
	binPlain, err := json.MarshalIndent(plain, "", "    ")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(binPlain))
}