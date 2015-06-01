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
