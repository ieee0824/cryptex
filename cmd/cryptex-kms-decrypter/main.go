package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws/credentials/stscreds"
	"github.com/aws/aws-sdk-go/aws/session"

	"github.com/ieee0824/cryptex"
	"github.com/ieee0824/cryptex/kms"
)

type option struct {
	inputFileName string
	text          string
}

func parseArgs() *option {
	ret := &option{}

	flag.StringVar(&ret.inputFileName, "i", "", "input file name")
	flag.StringVar(&ret.text, "t", "", "encrypted text")
	flag.Parse()
	if len(flag.Args()) != 0 && ret.text == "" {
		ret.text = flag.Args()[0]
	}

	return ret
}

func main() {
	opt := parseArgs()

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		AssumeRoleTokenProvider: stscreds.StdinTokenProvider,
		SharedConfigState:       session.SharedConfigEnable,
	}))

	kmsClient := kms.New(sess)
	c := cryptex.New(kmsClient)

	if opt.text != "" {
		container := &cryptex.Container{
			EncryptionType: "kms",
			Values:         opt.text,
		}
		row, err := c.Decrypt(container)
		if err != nil {
			panic(err)
		}
		json.NewEncoder(os.Stdout).Encode(row)

		return
	}

	switch opt.inputFileName {
	case "":
		var encryptStr string
		fmt.Scan(&encryptStr)
		container := &cryptex.Container{
			EncryptionType: "kms",
			Values:         encryptStr,
		}
		row, err := c.Decrypt(container)
		if err != nil {
			panic(err)
		}
		json.NewEncoder(os.Stdout).Encode(row)
	default:
	}
}
