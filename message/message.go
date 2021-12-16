package message

import (
	"errors"
	"github.com/mingzaily/go-wxcom"
)

// Message struct is used to compose and fire individual msg push from wxcom client.
type Message struct {
	client                 *wxcom.Client
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

// RespMessage struct holds response values of send msg.
type RespMessage struct {
	Errcode      int    `json:"errcode"`
	Errmsg       string `json:"errmsg"`
	Invaliduser  string `json:"invaliduser"`
	Invalidparty string `json:"invalidparty"`
	Invalidtag   string `json:"invalidtag"`
	Msgid        string `json:"msgid"`
	ResponseCode string `json:"response_code"`
}

// New method creates a new Message client with wxcom client.
func New(client *wxcom.Client) *Message {
	return &Message{
		client: client,
		path:   "/cgi-bin/msg/send",
	}
}

// Text method creates text msg.
func (m *Message) Text(content string) *text {
	return &text{
		base:    base{message: m},
		content: content,
	}
}

// Image method creates image msg.
func (m *Message) Image(mediaId string) *image {
	return &image{
		base:    base{message: m},
		mediaId: mediaId,
	}
}

// Voice method creates voice msg.
func (m *Message) Voice(mediaId string) *voice {
	return &voice{
		base:    base{message: m},
		mediaId: mediaId,
	}
}

// Video method creates video msg.
func (m *Message) Video(mediaId string) *video {
	return &video{
		base:    base{message: m},
		mediaId: mediaId,
	}
}

// File method creates file msg.
func (m *Message) File(mediaId string) *file {
	return &file{
		base:    base{message: m},
		mediaId: mediaId,
	}
}

// Textcard method creates textcard msg.
func (m *Message) Textcard(title, description, url string) *textcard {
	return &textcard{
		base:        base{message: m},
		title:       title,
		description: description,
		url:         url,
	}
}

// Markdown method creates markdown msg.
func (m *Message) Markdown(content string) *markdown {
	return &markdown{
		base:    base{message: m},
		content: content,
	}
}

func (m *Message) clone() *Message {
	newMessage := *m
	return &newMessage
}

func (m *Message) build() (map[string]interface{}, error) {

	if m.toUser == "" && m.toParty == "" && m.toTag == "" {
		return nil, errors.New("touser, toparty, totag cannot be empty at the same time")
	}

	body := map[string]interface{}{
		"touser":                   m.toUser,
		"toparty":                  m.toParty,
		"totag":                    m.toTag,
		"agentid":                  m.client.Agentid(),
		"enable_duplicate_check":   m.enableDuplicateCheck,
		"duplicate_check_interval": m.duplicateCheckInterval,
	}

	switch m.msgType {
	case "text":
		body["msgtype"] = "text"
		body["text"] = map[string]string{"content": m.content}
		body["safe"] = m.safe
		body["enable_id_trans"] = m.enableIdTrans
	case "image":
		body["msgtype"] = "image"
		body["image"] = map[string]string{"media_id": m.mediaId}
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
		body["textcard"] = map[string]string{"title": m.title, "description": m.description, "path": m.url, "btntxt": m.btnTxt}
		body["enable_id_trans"] = m.enableIdTrans
	case "markdown":
		body["msgtype"] = "markdown"
		body["markdown"] = map[string]string{"content": m.content}
	default:
		return nil, errors.New("unsupported msg type")
	}

	return body, nil
}

func (m *Message) send() (*RespMessage, error) {
	var response *RespMessage

	body, err := m.build()
	if err != nil {
		return nil, err
	}

	_, err = m.client.Resty.R().
		SetQueryParam("access_token", m.client.GetAccessToken()).
		SetHeader("Content-Type", "application/json; charset=UTF-8").
		SetBody(body).
		SetResult(&response).
		Post(m.path)
	if err != nil {
		return nil, err
	}

	if wxcom.IsTokenInvalidErr(response.Errcode, m.client) {
		return m.send()
	}

	return response, nil
}
