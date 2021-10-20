package wx_com

import (
	"testing"
)

func TestClient_GetAccessToken(t *testing.T) {
	c := New("", "", 0)

	accessToken := c.GetAccessToken()

	if accessToken == "" {
		t.Fatal("access_token empty")
	}

	if _, found := c.cache.Get("access_token"); !found {
		t.Fatal("access_token not cached")
	}
}
