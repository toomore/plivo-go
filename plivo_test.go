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
	TestDst      string = "123456789"
	TestSrc      string = "886976543210"
)

func TestMessage(t *testing.T) {
	account := &Account{User: TestUser, Password: TestPassword}
	now := time.Now()
	msg := NewMessage(TestDst, TestSrc,
		fmt.Sprintf("Hello 世界！%02d%02d%02d", now.Hour(), now.Minute(), now.Second()),
		account)
	log.Println(msg)
	msg.Send()
}

func ExampleMessage() {
	account := &Account{User: TestUser, Password: TestPassword}
	msg := NewMessage(TestDst, TestSrc, "Hello 世界！", account)
	fmt.Println(msg)
	msg.Send()
	// output:
	// <dst: "123456789" src: "886976543210" text: "Hello 世界！" Len: 9>
}
