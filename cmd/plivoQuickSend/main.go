// 快速單發 Plivo SMS.
//
/*
Install:

	go install github.com/toomore/plivo-go/cmd/plivoQuickSend

Usage:

	plivoQuickSend [flags]

The flags are:

	-src
		Plivo phone number.
	-dst
		Receiver phone number.
	-text
		SMS text content.
	-user
		API Auth ID
	-password
		API Auth Token

範例

	plivoQuickSend -src=12345678 -dst=886976543210 -text="測試簡訊" -user=xcdfasegg -password=oikjdndhy

*/
package main

import (
	"flag"
	"os"

	plivo "github.com/toomore/plivo-go"
)

var user = flag.String("user", "", "Plive User ID.")
var password = flag.String("password", "", "Plive User Token.")
var src = flag.String("src", "", "Plivo phone number.")
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
