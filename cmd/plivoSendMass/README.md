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
		Plive User ID.(Default get environment variable by the key named `PLIVOID`)
	-password
		Plive User Token.(Default get environment variable by the key named `PLIVOTOKEN`)

範例

	plivoSendMass -csv=/run/shm/list.csv -user=xcdfasegg -password=oikjdndhy

Docker
-------

Download image

    docker pull toomore/plivo-go

Send mass SMS with CSV.

    docker run -d -v <local_csv_path>:<container_csv_path> -e PLIVOID=<ID> -e PLIVOTOKEN=<TOKEN> -e PLIVOSRC=<SRC> toomore/plivo-go plivoSendMass -csv=<container_csv_path>

    549fc9920a6ec943a86b2f5afc569ef78d54a67e432e11888f0ddf5081158750

Read send logs

    docker logs 549fc9920a6ec943a86b2f5afc569ef78d54a67e432e11888f0ddf5081158750
