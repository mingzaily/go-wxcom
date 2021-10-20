package wxcom

// common
type appMessageCommon struct {
	Touser                 string `json:"touser"`
	Toparty                string `json:"toparty"`
	Totag                  string `json:"totag"`
	Msgtype                string `json:"msgtype"`
	Agentid                int    `json:"agentid"`
	Safe                   int    `json:"safe"`
	EnableIdTrans          int    `json:"enable_id_trans"`
	EnableDuplicateCheck   int    `json:"enable_duplicate_check"`
	DuplicateCheckInterval int    `json:"duplicate_check_interval"`
}

type appMessageContent struct {
	Content string `json:"content"`
}

// text message
type appTextMessageRequest struct {
	appMessageCommon
	Text appMessageContent `json:"text"`
}

func (a *appTextMessageRequest) sendable() bool {
	return !(a.Touser == "" && a.Toparty == "" && a.Totag == "")
}

func (a *appTextMessageRequest) setAgentid(agentid int) {
	a.Agentid = agentid
}

func NewAppTextMessageRequest(toUser, toParty, toTag, content string, safe int) *appTextMessageRequest {
	return &appTextMessageRequest{
		appMessageCommon: appMessageCommon{
			Touser:  toUser,
			Toparty: toParty,
			Totag:   toTag,
			Msgtype: "text",
			Safe:    safe,
		},
		Text: appMessageContent{
			Content: content,
		},
	}
}

// markdown message
type appMessageMarkdownRequest struct {
	appMessageCommon
	Markdown appMessageContent `json:"markdown"`
}

func (a *appMessageMarkdownRequest) sendable() bool {
	return !(a.Touser == "" && a.Toparty == "" && a.Totag == "")
}

func (a *appMessageMarkdownRequest) setAgentid(agentid int) {
	a.Agentid = agentid
}

func NewAppMarkdownMessageRequest(toUser, toParty, toTag, content string, safe int) *appMessageMarkdownRequest {
	return &appMessageMarkdownRequest{
		appMessageCommon: appMessageCommon{
			Touser:  toUser,
			Toparty: toParty,
			Totag:   toTag,
			Msgtype: "markdown",
			Safe:    safe,
		},
		Markdown: appMessageContent{
			Content: content,
		},
	}
}

type AppMessageResponse struct {
	Errcode      int    `json:"errcode"`
	Errmsg       string `json:"errmsg"`
	Invaliduser  string `json:"invaliduser"`
	Invalidparty string `json:"invalidparty"`
	Invalidtag   string `json:"invalidtag"`
	Msgid        string `json:"msgid"`
	ResponseCode string `json:"response_code"`
}
