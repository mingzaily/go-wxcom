package message

import (
	"github.com/mingzaily/wxcom-sdk"
	"testing"
)

func TestMessage_SendAppTxtMessage(t *testing.T) {
	client := wx_com.New("", "", 0)
	m := NewWithClient(client)
	r := NewAppTextMessageRequest("", "", "", "", 0)

	_, err := m.SendAppMessage(r)
	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestMessage_SendAppMarkdownMessage(t *testing.T) {
	client := wx_com.New("", "", 1000195).SetDebug(true)
	m := NewWithClient(client)
	r := NewAppMarkdownMessageRequest("", "", "", "", 0)

	_, err := m.SendAppMessage(r)
	if err != nil {
		t.Fatal(err.Error())
	}
}
