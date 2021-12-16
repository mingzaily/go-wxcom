package message_test

import (
	"github.com/mingzaily/go-wxcom"
	"github.com/mingzaily/go-wxcom/message"
	"reflect"
	"testing"
)

var msg = message.NewWithApp(wxcom.New("12345", "123456", 123))

func TestMessage_NotToPeople(t *testing.T) {
	_, err := msg.Text("你的快递已到，请携带工卡前往邮件中心领取。").Send()

	assertEqual(t, err.Error(), "toUser, toParty, toTag cannot be empty at the same time")
}

func TestMessage_ToUser(t *testing.T) {
	m := msg.ToUser([]string{"test"}).Text("你的快递已到，请携带工卡前往邮件中心领取。")

	assertEqual(t, m.ToJson(), "{\"agentid\":123,\"enable_id_trans\":0,\"msgtype\":\"text\",\"safe\":0,\"text\":{\"content\":\"你的快递已到，请携带工卡前往邮件中心领取。\"},\"touser\":\"test\"}")
}

func TestMessage_ToParty(t *testing.T) {
	m := msg.ToParty([]string{"test"}).Text("你的快递已到，请携带工卡前往邮件中心领取。")

	assertEqual(t, m.ToJson(), "{\"agentid\":123,\"enable_id_trans\":0,\"msgtype\":\"text\",\"safe\":0,\"text\":{\"content\":\"你的快递已到，请携带工卡前往邮件中心领取。\"},\"toparty\":\"test\"}")
}

func TestMessage_ToTag(t *testing.T) {
	m := msg.ToTag([]string{"test"}).Text("你的快递已到，请携带工卡前往邮件中心领取。")

	assertEqual(t, m.ToJson(), "{\"agentid\":123,\"enable_id_trans\":0,\"msgtype\":\"text\",\"safe\":0,\"text\":{\"content\":\"你的快递已到，请携带工卡前往邮件中心领取。\"},\"totag\":\"test\"}")
}

func TestMessage_SetSafe(t *testing.T) {
	m := msg.ToUser([]string{"test"}).SetSafe(1).Text("你的快递已到，请携带工卡前往邮件中心领取。")

	assertEqual(t, m.ToJson(), "{\"agentid\":123,\"enable_id_trans\":0,\"msgtype\":\"text\",\"safe\":1,\"text\":{\"content\":\"你的快递已到，请携带工卡前往邮件中心领取。\"},\"touser\":\"test\"}")
}

func assertEqual(t *testing.T, e, g interface{}) (r bool) {
	if !equal(e, g) {
		t.Errorf("Expected [%v], got [%v]", e, g)
	}

	return true
}

func assertNotEqual(t *testing.T, e, g interface{}) (r bool) {
	if equal(e, g) {
		t.Errorf("Expected [%v], got [%v]", e, g)
	}

	return true
}

func equal(expected, got interface{}) bool {
	return reflect.DeepEqual(expected, got)
}
