package wxcom

import (
	"encoding/json"
	"errors"
	"log"
	"strings"
)

// Message struct is used to compose and fire individual message push with wxcom Wxcom client.
type Message struct {
	wx                     *Wxcom
	path                   string
	msgType                string
	toUser                 string
	toParty                string
	toTag                  string
	safe                   int
	content                string
	mediaId                string
	title                  string
	description            string
	musicUrl               string
	hpMusicUrl             string
	enableIdTrans          int
	url                    string
	btnTxt                 string
	enableDuplicateCheck   int
	duplicateCheckInterval int
}

// RespMessage struct holds response values of send message.
type RespMessage struct {
	Errcode      int    `json:"errcode"`
	Errmsg       string `json:"errmsg"`
	Invaliduser  string `json:"invaliduser"`
	Invalidparty string `json:"invalidparty"`
	Invalidtag   string `json:"invalidtag"`
	Msgid        string `json:"msgid"`
	ResponseCode string `json:"response_code"`
}

// ToUser method sets to user to in the current message.
func (m *Message) ToUser(userList []string) *Message {
	m.toUser = strings.Join(userList, "|")
	return m
}

// ToParty method sets to party to in the current message.
func (m *Message) ToParty(partyList []string) *Message {
	m.toParty = strings.Join(partyList, "|")
	return m
}

// ToTag method sets to tag to in the current message.
func (m *Message) ToTag(tagList []string) *Message {
	m.toTag = strings.Join(tagList, "|")
	return m
}

// DuplicateCheck method enables the duplicate check.
// Param example:
// 0, 0 duplicate check is not enabled.
// 1, 1800 duplicate check within 1800 seconds.
func (m *Message) DuplicateCheck(enableDuplicateCheck, duplicateCheckInterval int) *Message {
	m.enableDuplicateCheck = enableDuplicateCheck
	if m.enableDuplicateCheck == 1 {
		m.duplicateCheckInterval = duplicateCheckInterval
	}
	return m
}

// Clone method create the new message wx.
func (m *Message) Clone() *Message {
	return m.clone()
}

// clone method create the new message wx.
func (m *Message) clone() *Message {
	newMessage := *m
	return &newMessage
}

// genRequestParam method generate http request params.
func (m *Message) genRequestParam() (map[string]interface{}, error) {

	if m.toUser == "" && m.toParty == "" && m.toTag == "" {
		return nil, errors.New("toUser, toParty, toTag cannot be empty at the same time")
	}

	body := map[string]interface{}{
		"agentid": m.wx.GetAgentid(),
	}
	if m.toUser != "" {
		body["touser"] = m.toUser
	}
	if m.toParty != "" {
		body["toparty"] = m.toParty
	}
	if m.toTag != "" {
		body["totag"] = m.toTag
	}
	if m.enableDuplicateCheck != 0 {
		body["enable_duplicate_check"] = m.enableDuplicateCheck
	}
	if m.enableDuplicateCheck != 0 {
		body["duplicate_check_interval"] = m.duplicateCheckInterval
	}

	switch m.msgType {
	case "text":
		body["text"] = map[string]string{"content": m.content}
		body["safe"] = m.safe
		body["enable_id_trans"] = m.enableIdTrans
	case "image":
		body["msgtype"] = "image"
		body["image"] = map[string]string{"media_id": m.mediaId}
		body["safe"] = m.safe
	case "voice":
		body["msgtype"] = "voice"
		body["voice"] = map[string]string{"media_id": m.mediaId}
	case "video":
		body["msgtype"] = "video"
		body["video"] = map[string]string{"media_id": m.mediaId, "title": m.title, "description": m.description}
		body["safe"] = m.safe
	case "file":
		body["msgtype"] = "file"
		body["file"] = map[string]string{"media_id": m.mediaId}
		body["safe"] = m.safe
	case "textcard":
		body["msgtype"] = "textcard"
		body["textcard"] = map[string]string{"title": m.title, "description": m.description, "url": m.url, "btntxt": m.btnTxt}
		body["enable_id_trans"] = m.enableIdTrans
	case "markdown":
		body["msgtype"] = "markdown"
		body["markdown"] = map[string]string{"content": m.content}
	default:
		return nil, errors.New("unsupported msg type")
	}

	body["msgtype"] = m.msgType

	return body, nil
}

// toJson method return message string.
func (m *Message) toJson() string {
	param, err := m.genRequestParam()
	if err != nil {
		return ""
	}

	paramBytes, err := json.Marshal(param)
	if err != nil {
		log.Fatalln(err)
		return ""
	}

	return string(paramBytes)
}

// send method does send message.
func (m *Message) send() (*RespMessage, error) {
	var response *RespMessage

	body, err := m.genRequestParam()
	if err != nil {
		return nil, err
	}

	_, err = m.wx.Resty.R().
		SetQueryParam("access_token", m.wx.GetAccessToken()).
		SetHeader("Content-Type", "application/json; charset=UTF-8").
		SetBody(body).
		SetResult(&response).
		SetError(&response).
		Post(m.path)
	if err != nil {
		return nil, err
	}

	if m.wx.isTokenInvalidErr(response.Errcode) {
		return m.send()
	}

	return response, nil
}

// Text method creates text message.
func (m *Message) Text(content string) *text {
	return &text{
		message: m,
		content: content,
	}
}

// Image method creates image message.
func (m *Message) Image(mediaId string) *image {
	return &image{
		message: m,
		mediaId: mediaId,
	}
}

// Voice method creates voice message.
func (m *Message) Voice(mediaId string) *voice {
	return &voice{
		message: m,
		mediaId: mediaId,
	}
}

// Video method creates video message.
func (m *Message) Video(mediaId string) *video {
	return &video{
		message: m,
		mediaId: mediaId,
	}
}

// File method creates file message.
func (m *Message) File(mediaId string) *file {
	return &file{
		message: m,
		mediaId: mediaId,
	}
}

// Textcard method creates textcard message.
func (m *Message) Textcard(title, description, url string) *textcard {
	return &textcard{
		message:     m,
		title:       title,
		description: description,
		url:         url,
	}
}

// Markdown method creates markdown message.
func (m *Message) Markdown(content string) *markdown {
	return &markdown{
		message: m,
		content: content,
	}
}
