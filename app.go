package wxcom

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/patrickmn/go-cache"
	"time"
)

// App struct is used to create wxcom client.
//
// The cache uses patrickmn/go-cache.
// You can refer to related documents(https://github.com/patrickmn/go-cache) if necessary.
//
// The resty uses go-resty/resty/v2.
// You can refer to related documents(https://github.com/go-resty/resty) if necessary.
type App struct {
	corpid     string
	corpsecret string
	agentid    int
	cache      *cache.Cache
	Resty      *resty.Client
}

// New method creates a new wxcom App client.
func New(corpid, corpsecret string, agentid int) *App {
	return &App{
		corpid:     corpid,
		corpsecret: corpsecret,
		agentid:    agentid,
		cache:      cache.New(5*time.Minute, 10*time.Minute),
		Resty:      resty.New().SetHostURL("https://qyapi.weixin.qq.com/"),
	}
}

// RestyDebug method enables the debug mode on Resty client. App logs details of every request and response.
// For `Request` it logs information such as HTTP verb, Relative URL path, Host, Headers, Body if it has one.
// For `Response` it logs information such as Status, Response Time, Headers, Body if it has one.
func (a *App) RestyDebug(d bool) *App {
	a.Resty.SetDebug(d)
	return a
}

// GetAgentid method get agentid from client
func (a *App) GetAgentid() int {
	return a.agentid
}

//‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾
// Common method
//__________________________________________________

// IsTokenInvalidErr method check whether the token has expired.
func IsTokenInvalidErr(errcode int, a *App) bool {
	if errcode == 42001 || errcode == 40014 {
		a.cache.Delete("access_token_" + fmt.Sprintf("%d", a.agentid))
		return true
	}
	return false
}
