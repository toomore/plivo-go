// 大量發送 Plivo SMS.
//
/*
Install:

	go install github.com/toomore/plivo-go/cmd/plivoSendMass

Usage:

	plivoSendMass [flags]

The flags are:

	-csv
		CSV 檔案位置，需包含三個欄位 `dst` 收件人電話（8869xxxxxxxx）, `src` 寄件者電話（Plivo 申請的電話）, `text` 簡訊內容（支援長內容）
	-user
		API Auth ID
	-password
		API Auth Token

範例

	plivoSendMass -csv=/run/shm/list.csv -user=xcdfasegg -password=oikjdndhy

*/
package main

import (
	"encoding/csv"
	"flag"
	"log"
	"os"
	"runtime"
	"sync"

	plivo "github.com/toomore/plivo-go"
)

var user = flag.String("user", "", "Plive User ID.")
var password = flag.String("password", "", "Plive User Token.")
var csvpath = flag.String("csv", "", "CSV file path.")
var wg sync.WaitGroup

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	flag.Parse()

	if flag.NFlag() == 0 {
		flag.PrintDefaults()
		os.Exit(0)
	}

	var account = &plivo.Account{User: *user, Password: *password}

	data, _ := os.Open(*csvpath)
	defer data.Close()

	alldata, _ := csv.NewReader(data).ReadAll()
	for _, v := range alldata[1:] {
		wg.Add(1)
		go func(v []string) {
			defer wg.Done()
			runtime.Gosched()

			value := make(map[string]string)
			for i, key := range alldata[0] {
				value[key] = v[i]
			}

			msg := plivo.NewMessage(value["dst"], value["src"], value["text"], account)
			log.Printf("Start %s", msg)
			msg.Send()
		}(v)
	}
	wg.Wait()
}
