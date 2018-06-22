package main

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials/stscreds"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/ieee0824/cryptex"
	"github.com/ieee0824/cryptex/kms"
	"github.com/ieee0824/getenv"
	"github.com/joho/godotenv"
)

var (
	KMS_KEY_ID string
)

func init() {
	godotenv.Load(".env")
	KMS_KEY_ID = getenv.String("KMS_KEY_ID")
}

func main() {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		AssumeRoleTokenProvider: stscreds.StdinTokenProvider,
		SharedConfigState:       session.SharedConfigEnable,
	}))
	sess.Config.Region = aws.String(getenv.String("AWS_REGION"))

	kmsClient := kms.New(sess)

	kmsClient.SetKey(KMS_KEY_ID)

	plainMap := map[string]interface{}{
		"hoge":  "huga",
		"foo":   "bar",
		"int":   0,
		"float": 1.1,
		"sub_map": map[string]interface{}{
			"alice": 12,
			"bob":   25,
		},
		"very_long_string": "$$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy$0fkqJkYTy0fkqJkYTy",
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

	p, err := c.Decrypt(cipher)
	if err != nil {
		panic(err)
	}
	fmt.Println(reflect.DeepEqual(plainMap, p))
}
