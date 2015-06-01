// 快速單發 Plivo SMS.
//
/*
Install:

	go install github.com/toomore/plivo-go/cmd/plivoQuickSend

Usage:

	plivoQuickSend [flags]

The flags are:

	-src
		Plivo phone number.(Default get environment variable by the key named `PLIVOSRC`)
	-dst
		Receiver phone number.
	-text
		SMS text content.
	-user
		API Auth ID.(Default get environment variable by the key named `PLIVOID`)
	-password
		API Auth Token.(Default get environment variable by the key named `PLIVOTOKEN`)

範例

	plivoQuickSend -src=12345678 -dst=886976543210 -text="測試簡訊" -user=xcdfasegg -password=oikjdndhy

*/
package main

import (
	"flag"
	"os"

	plivo "github.com/toomore/plivo-go"
)

var user = flag.String("user", os.Getenv("PLIVOID"), "Plive User ID.(Default get environment variable by the key named `PLIVOID`)")
var password = flag.String("password", os.Getenv("PLIVOTOKEN"), "Plive User Token.(Default get environment variable by the key named `PLIVOTOKEN`)")
var src = flag.String("src", os.Getenv("PLIVOSRC"), "Plivo phone number.(Default get environment variable by the key named `PLIVOSRC`)")
var dst = flag.String("dst", "", "Receiver phone number.")
var text = flag.String("text", "", "SMS text content.")

func main() {
	flag.Parse()

	if flag.NFlag() == 0 {
		flag.PrintDefaults()
		os.Exit(0)
	}

	msg := plivo.NewMessage(*dst, *src, *text, &plivo.Account{User: *user, Password: *password})
	msg.Send()
}
