package main

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go/aws/credentials/stscreds"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/ieee0824/cryptex"
	"github.com/ieee0824/cryptex/kms"
)

func main() {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		AssumeRoleTokenProvider: stscreds.StdinTokenProvider,
		SharedConfigState:       session.SharedConfigEnable,
	}))

	kmsClient := kms.New(sess)

	kmsClient.SetKey("kms key id")

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

	c := cryptex.New(kmsClient)

	cipher, err := c.Encrypt(plainMap)
	if err != nil {
		panic(err)
	}

	bin, err := json.MarshalIndent(cipher, "", "    ")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(bin))
}
