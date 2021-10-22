package wxcom

import (
	"os"
	"reflect"
	"strconv"
	"testing"
)

var agentid, _ = strconv.Atoi(os.Getenv("AGENTID"))
var corpid = os.Getenv("CORPID")
var corpsecret = os.Getenv("CORPSECRET")

var wx = New(corpid, corpsecret, agentid)

func TestWxcom_SetRestyDebug(t *testing.T) {
	wx.SetRestyDebug(true)

	assertEqual(t, wx.resty.Debug, true)

	wx.SetRestyDebug(false)

	assertEqual(t, wx.resty.Debug, false)
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
