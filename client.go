package wx_com

import (
	"github.com/go-resty/resty/v2"
	"github.com/patrickmn/go-cache"
	"sync"
	"time"
)

type Client struct {
	corpid     string
	corpsecret string
	agentid    int
	cache      *cache.Cache
	http       *resty.Client
	mutex      *sync.RWMutex
}

func New(corpid, corpsecret string, agentid int) *Client {
	return &Client{
		corpid:     corpid,
		corpsecret: corpsecret,
		agentid:    agentid,
		cache:      cache.New(5*time.Minute, 10*time.Minute),
		http:       resty.New(),
		mutex:      &sync.RWMutex{},
	}
}

func (c *Client) C() *resty.Client {
	return c.http
}

func (c *Client) SetDebug(d bool) *Client {
	c.http.SetDebug(d)
	return c
}

func (c *Client) SetCorpid(d string) *Client {
	c.corpid = d
	return c
}

func (c *Client) SetCorpsecret(d string) *Client {
	c.corpsecret = d
	return c
}

func (c *Client) SetAgentid(d int) *Client {
	c.agentid = d
	return c
}

func (c *Client) GetAgentid() int {
	return c.agentid
}
