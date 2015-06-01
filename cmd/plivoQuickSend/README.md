plivoQuickSend
===============

[![GoDoc](https://godoc.org/github.com/toomore/plivo-go/cmd/plivoQuickSend?status.svg)](https://godoc.org/github.com/toomore/plivo-go/cmd/plivoQuickSend)

大量發送簡訊

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

