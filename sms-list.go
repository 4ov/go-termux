package termux

import (
	"encoding/json"
	"fmt"
)

const (
	SmsTypeAll = iota
	SmsTypeInbox
	SmsTypeSent
	SmsTypeDrafts
	SmsTypeOutbox
)

type Sms struct {
	ThreadId int    `json:"threadid"`
	Type     string `json:"type"`
	Read     bool   `json:"read"`
	Number   string `json:"number"`
	Recived  string `json:"received"`
	Body     string `json:"body"`
	Id       int    `json:"_id"`
}

type SmsOptions struct {
	// offset in sms list (default: 10)
	Limit int `arg:"-l"`
	// offset in sms list (default: 0)
	Offset int `arg:"-o"`
	// the type of messages to list (default: all):
	// all|inbox|sent|draft|outbox
	Type int `arg:"-t"`
	// (unique item per conversation)
	ConversationList bool `arg:"-c"`
	// the number for locate messages
	From string `arg:"-f"`
	// (obsolete) show phone numbers
	ShowNumbers bool `arg:"-n"`
	// (obsolete) show dates when messages were created
	ShowDates bool `arg:"-d"`
}

func SmsList(options SmsOptions) ([]Sms, error) {
	args := ReadyArgs(options)
	output, err := CallCommand("termux-sms-list", args...)
	if err != nil {
		fmt.Println(err)
		return ([]Sms{}), err
	}
	result := []Sms{}
	err = json.Unmarshal(output, &result)
	if err != nil {
		fmt.Println(err)
		return ([]Sms{}), err
	}
	return result, nil
}
