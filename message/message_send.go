package message

import (
	"strings"
)

type SendMessage interface {
	Send() (*RespMessage, error)
}

type base struct {
	message                *Message
	msgType                string
	toUser                 string
	toParty                string
	toTag                  string
	safe                   int
	enableIdTrans          int
	enableDuplicateCheck   int
	duplicateCheckInterval int
}

func (b *base) build() *Message {
	nm := b.message.clone()
	nm.toUser = b.toUser
	nm.toParty = b.toParty
	nm.toTag = b.toTag
	return nm
}

// ToUser method sets to user to in the current msg.
func (b *base) ToUser(userList []string) *base {
	b.toUser = strings.Join(userList, "|")
	return b
}

// ToParty method sets to party to in the current msg.
func (b *base) ToParty(partyList []string) *base {
	b.toParty = strings.Join(partyList, "|")
	return b
}

// ToTag method sets to tag to in the current msg.
func (b *base) ToTag(tagList []string) *base {
	b.toTag = strings.Join(tagList, "|")
	return b
}

// Safe method sets the msg is confident.
func (b *base) Safe(safe int) *base {
	b.safe = safe
	return b
}

// EnableIdTrans method sets the msg enable id translation.
func (b *base) EnableIdTrans(enableIdTrans int) *base {
	b.enableIdTrans = enableIdTrans
	return b
}

// DuplicateCheck method enables the duplicate check.
// Param example:
// 0, 0 duplicate check is not enabled.
// 1, 1800 duplicate check within 1800 seconds.
func (b *base) DuplicateCheck(enableDuplicateCheck, duplicateCheckInterval int) *base {
	b.enableDuplicateCheck = enableDuplicateCheck
	if b.enableDuplicateCheck == 1 {
		b.duplicateCheckInterval = duplicateCheckInterval
	}
	return b
}

// text struct is used to compose text msg push from msg client.
type text struct {
	base
	content string
}

// Send method does Send text msg.
func (t *text) Send() (*RespMessage, error) {
	msg := t.build()
	msg.msgType = "text"
	msg.content = t.content
	return msg.send()
}

// image struct is used to compose image msg push from msg client.
type image struct {
	base
	mediaId string
}

// Send method does Send image image msg.
func (i *image) Send() (*RespMessage, error) {
	msg := i.build()
	msg.msgType = "image"
	msg.mediaId = i.mediaId
	return msg.send()
}

// voice struct is used to compose voice msg push from msg client.
type voice struct {
	base
	mediaId string
}

// Send method does Send voice msg.
func (v *voice) Send() (*RespMessage, error) {
	msg := v.build()
	msg.msgType = "voice"
	msg.mediaId = v.mediaId
	return msg.send()
}

// video struct is used to compose video msg push from msg client.
type video struct {
	base
	mediaId     string
	title       string
	description string
}

// Title method sets the video msg title.
func (v *video) Title(title string) *video {
	v.title = title
	return v
}

// Description method sets the video msg description.
func (v *video) Description(description string) *video {
	v.description = description
	return v
}

// Send method does Send video msg.
func (v *video) Send() (*RespMessage, error) {
	msg := v.build()
	msg.msgType = "video"
	msg.mediaId = v.mediaId
	msg.title = v.title
	msg.description = v.description
	return msg.send()
}

// file struct is used to compose file msg push from msg client.
type file struct {
	base
	mediaId string
}

// Send method does Send file msg.
func (f *file) Send() (*RespMessage, error) {
	msg := f.build()
	msg.mediaId = f.mediaId
	return msg.send()
}

// textcard struct is used to compose textcard msg push from msg client.wx_message
type textcard struct {
	base
	title       string
	description string
	url         string
	btnTxt      string
}

// Btntxt method sets the textcard msg btn txt.
func (t *textcard) Btntxt(btnTxt string) *textcard {
	t.btnTxt = btnTxt
	return t
}

// Send method does Send textcard msg.
func (t *textcard) Send() (*RespMessage, error) {
	msg := t.build()
	msg.title = t.title
	msg.description = t.description
	msg.url = t.url
	msg.btnTxt = t.btnTxt
	return msg.send()
}

// markdown struct is used to compose markdown msg push from msg client.
type markdown struct {
	base
	content string
}

// Send method does Send markdown msg.
func (m *markdown) Send() (*RespMessage, error) {
	msg := m.build()
	msg.content = msg.mediaId
	return msg.send()
}
