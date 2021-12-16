package wxcom

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/patrickmn/go-cache"
	"time"
)

// Client struct is used to create wxcom client.
//
// The cache uses patrickmn/go-cache.
// You can refer to related documents(https://github.com/patrickmn/go-cache) if necessary.
//
// The resty uses go-resty/resty/v2.
// You can refer to related documents(https://github.com/go-resty/resty) if necessary.
type Client struct {
	corpid     string
	corpsecret string
	agentid    int
	cache      *cache.Cache
	Resty      *resty.Client
}

// New method creates a new wxcom client.
func New(corpid, corpsecret string, agentid int) *Client {
	return &Client{
		corpid:     corpid,
		corpsecret: corpsecret,
		agentid:    agentid,
		cache:      cache.New(5*time.Minute, 10*time.Minute),
		Resty:      resty.New().SetHostURL("https://qyapi.weixin.qq.com/"),
	}
}

// RestyDebug method enables the debug mode on Resty client. Client logs details of every request and response.
// For `Request` it logs information such as HTTP verb, Relative URL path, Host, Headers, Body if it has one.
// For `Response` it logs information such as Status, Response Time, Headers, Body if it has one.
func (c *Client) RestyDebug(d bool) *Client {
	c.Resty.SetDebug(d)
	return c
}

// Agentid method get agentid from client
func (c *Client) Agentid() int {
	return c.agentid
}

//‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾
// Common method
//__________________________________________________

// IsTokenInvalidErr method check whether the token has expired.
func IsTokenInvalidErr(errcode int, w *Client) bool {
	if errcode == 42001 || errcode == 40014 {
		w.cache.Delete("access_token_" + fmt.Sprintf("%d", w.agentid))
		return true
	}
	return false
}
