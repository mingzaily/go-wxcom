package wxcom_test

import (
	"github.com/mingzaily/go-wxcom"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

var wx = wxcom.New("123", "321", 1)

func TestWxcom_GetAgentid(t *testing.T) {
	assertEqual(t, wx.GetAgentid(), 1)
}

func TestWxcom_GetAccessToken(t *testing.T) {
	ts := createTestServer(t)
	defer ts.Close()

	tempWx := wxcom.New("123", "321", 123)
	tempWx.Resty.SetBaseURL(ts.URL)

	// from server
	assertEqual(t, tempWx.GetAccessToken(), "token")
	// form token
	assertEqual(t, tempWx.GetAccessToken(), "token")
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

func createTestServer(t *testing.T) *httptest.Server {
	// for test invalid access token, retry two time.
	time := 0

	fn := func(w http.ResponseWriter, r *http.Request) {
		t.Logf("Method: %v", r.Method)
		t.Logf("Path: %v", r.URL.Path)

		switch r.URL.Path {
		case "/cgi-bin/gettoken":
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte("{\"errcode\":0,\"errmsg\":\"ok\",\"access_token\":\"token\",\"expires_in\":7200}"))
		case "/cgi-bin/message/send":
			w.Header().Set("Content-Type", "application/json")
			if time > 0 {
				_, _ = w.Write([]byte("{\"errcode\":0,\"errmsg\":\"ok\",\"msgid\":\"msgid\"}"))
			} else {
				_, _ = w.Write([]byte("{\"errcode\":42001,\"errmsg\":\"invalid access_token\"}"))
			}
			time++
		case "/cgi-bin/user/getuserinfo":
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte("{\"errcode\":0,\"errmsg\":\"ok\",\"UserId\":\"test_user\",\"DeviceId\":\"device\"}"))
		}
	}

	return httptest.NewServer(http.HandlerFunc(fn))
}
