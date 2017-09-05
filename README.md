# cryptex
encrypt map, only value.

# Example

## Use KMS

```
sess := session.Must(session.NewSessionWithOptions(session.Options{
	AssumeRoleTokenProvider: stscreds.StdinTokenProvider,
	SharedConfigState:       session.SharedConfigEnable,
}))

kmsClient := kms.New(sess)

kmsClient.SetKey("kms key id")

plainMap := map[string]interface{}{
	"hoge": "huga",
	"foo": "bar",
	"int": 0,
	"float": 1.1,
	"sub_map": map[string]interface{}{
		"alice": 12,
		"bob": 25,
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
```

```
$ go run examples/kms/main.go 
{
    "float": "AQICAHhHV0+8t79k1rzbJjVWp5OdYcOSrGZYstS+b9s5iJx6qAEVHYE2BmWZDCdjzBsp8BrUAAAAazBpBgkqhkiG9w0BBwagXDBaAgEAMFUGCSqGSIb3DQEHATAeBglghkgBZQMEAS4wEQQM1s8D/Z0xT4tzl4vBAgEQgCjF5DG4u+ta4G0hPppxKs/MvEnRKJWsToRelE70RBD2SHpQOM3HoHlE",
    "foo": "AQICAHhHV0+8t79k1rzbJjVWp5OdYcOSrGZYstS+b9s5iJx6qAFvrN28mE5e8hYBd9QFOiBcAAAAbTBrBgkqhkiG9w0BBwagXjBcAgEAMFcGCSqGSIb3DQEHATAeBglghkgBZQMEAS4wEQQMWe4H/D6jLl4yoE6MAgEQgCopODL6ZdLg+QEIL3Jt5I1iIu5EZssAS9ThFdaQGM91omzvp5oZOTjStTc=",
    "hoge": "AQICAHhHV0+8t79k1rzbJjVWp5OdYcOSrGZYstS+b9s5iJx6qAETui2e7OSBvojVQ/oinP1HAAAAbjBsBgkqhkiG9w0BBwagXzBdAgEAMFgGCSqGSIb3DQEHATAeBglghkgBZQMEAS4wEQQMaY5FsQaXGrdYygevAgEQgCs1n311741Wp3jEvTvPE+TMjRPiwjBAWi6QgWAELt2cq2n+7wP25b+hI6dB",
    "int": "AQICAHhHV0+8t79k1rzbJjVWp5OdYcOSrGZYstS+b9s5iJx6qAHfA/cMHMgIA03TYEA2mUIEAAAAaTBnBgkqhkiG9w0BBwagWjBYAgEAMFMGCSqGSIb3DQEHATAeBglghkgBZQMEAS4wEQQMswEQ5Jg316jaBMWAAgEQgCZ2py5xmFUop9IC0Q9+nTrMVbdjSfCuU95oGSTW5JM/zmEBvQAvnw==",
    "sub_map": {
        "alice": "AQICAHhHV0+8t79k1rzbJjVWp5OdYcOSrGZYstS+b9s5iJx6qAFtQ+9LqI93fp8UuqoOl87UAAAAajBoBgkqhkiG9w0BBwagWzBZAgEAMFQGCSqGSIb3DQEHATAeBglghkgBZQMEAS4wEQQM3jsNIGA2gHx9gGw5AgEQgCepnkvmAdfQRvB7d8fW64719oz9A8VDOld/Cwzg7alUw+E/cJNqKlI=",
        "bob": "AQICAHhHV0+8t79k1rzbJjVWp5OdYcOSrGZYstS+b9s5iJx6qAGMkacS1QStfuHry/upq9ZwAAAAajBoBgkqhkiG9w0BBwagWzBZAgEAMFQGCSqGSIb3DQEHATAeBglghkgBZQMEAS4wEQQMgl1jcrTJEYhr6He+AgEQgCcM26QGSFW8F8bF2FIG8W3z8GubAVYh3vPz8+/FhWI42zewPh4x2jM="
    }
}

```
