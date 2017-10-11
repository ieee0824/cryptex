package main

import (
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go/aws/credentials/stscreds"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/ieee0824/cryptex"
	"github.com/ieee0824/cryptex/kms"
	"github.com/joho/godotenv"
	"github.com/ieee0824/getenv"
	"github.com/aws/aws-sdk-go/aws"
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
