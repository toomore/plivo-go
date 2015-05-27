package plivo

import (
	"fmt"
	"log"
	"testing"
	"time"
)

const (
	TestUser     string = ""
	TestPassword string = ""
	TestDst      string = ""
	TestSrc      string = ""
)

func TestMessage(t *testing.T) {
	account := &Account{User: TestUser, Password: TestPassword}
	now := time.Now()
	m := NewMessage(TestDst, TestSrc,
		fmt.Sprintf("Hello 世界！%02d%02d%02d", now.Hour(), now.Minute(), now.Second()),
		account)
	log.Println(m)
	m.Send()
}
