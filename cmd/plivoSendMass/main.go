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

		value := make(map[string]string)
		for keyi, key := range alldata[0] {
			value[key] = v[keyi]
		}

		go func(value map[string]string) {
			log.Println(value)
			runtime.Gosched()
			msg := plivo.NewMessage(value["dst"], value["src"], value["text"], account)
			msg.Send()
			defer wg.Done()
		}(value)
	}
	wg.Wait()
}
