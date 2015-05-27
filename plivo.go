// Package plivo is to send SMS tools.
package plivo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/url"
	"path"
	"strings"
	"time"
	"unicode/utf8"
)

// PlivoAPI URL.
const PlivoAPI = "https://api.plivo.com/v1/Account"

// Fixed http too many open files.
var httpClient = &http.Client{Transport: &http.Transport{
	Proxy: http.ProxyFromEnvironment,
	Dial: (&net.Dialer{
		Timeout:   0,
		KeepAlive: 0,
	}).Dial,
	TLSHandshakeTimeout: 10 * time.Second,
},
}

// Message struct
type Message struct {
	dst     string
	src     string
	text    string
	account *Account
}

// NewMessage new a Message.
func NewMessage(dst, src, text string, account *Account) *Message {
	return &Message{dst: dst, src: src, text: text, account: account}
}

func (m Message) toMap() map[string]string {
	var result = make(map[string]string)
	result["dst"] = m.dst
	result["src"] = m.src
	result["text"] = m.text
	return result
}

// Send to send SMS.
func (m Message) Send() {
	jsonData, _ := json.Marshal(m.toMap())
	req, _ := http.NewRequest("POST", m.URL("/Message/"), bytes.NewReader(jsonData))
	req.URL.User = url.UserPassword(m.account.User, m.account.Password)
	req.Header = http.Header{"Content-Type": {"application/json"}}
	if resp, err := httpClient.Do(req); err == nil {
		if body, err := ioutil.ReadAll(resp.Body); err == nil {
			defer resp.Body.Close()
			log.Printf("%s\n", body)
		}
	}
}

// String to print string.
func (m Message) String() string {
	return fmt.Sprintf(`<dst: "%s" src: "%s" text: "%s" Len: %d>`, m.dst, m.src, m.text, m.Len())
}

// Len to show length of Message text.
func (m Message) Len() int {
	return utf8.RuneCountInString(m.text)
}

// URL to render api full URL.
func (m Message) URL(urlpath string) string {
	URLPath, _ := url.ParseRequestURI(PlivoAPI)
	URLPath.Path = path.Join(URLPath.Path, m.account.User, urlpath)
	if strings.LastIndex(urlpath, "/") >= 0 {
		return fmt.Sprintf("%s/", URLPath.String())
	}
	return URLPath.String()
}

// Account struct
type Account struct {
	User     string
	Password string
}
