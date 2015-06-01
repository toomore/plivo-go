plivoSendMass
==============

[![GoDoc](https://godoc.org/github.com/toomore/plivo-go/cmd/plivoSendMass?status.svg)](https://godoc.org/github.com/toomore/plivo-go/cmd/plivoSendMass)

大量發送簡訊

Install:

	go install github.com/toomore/plivo-go/cmd/plivoSendMass

Usage:

	plivoSendMass [flags]

The flags are:

	-csv
		CSV 檔案位置，需包含三個欄位 `dst` 收件人電話（8869xxxxxxxx）, `src` 寄件者電話（Plivo 申請的電話）, `text` 簡訊內容（支援長內容）
	-src
		Plivo phone number.(Default get environment variable by the key named `PLIVOSRC`)
	-user
		API Auth ID
	-password
		API Auth Token

範例

	plivoSendMass -csv=/run/shm/list.csv -user=xcdfasegg -password=oikjdndhy

