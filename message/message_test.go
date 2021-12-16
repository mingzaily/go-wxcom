package message_test

import (
	"github.com/mingzaily/go-wxcom"
	"github.com/mingzaily/go-wxcom/message"
	"reflect"
	"testing"
)

var msg = message.New(wxcom.New("", "", 0))

func TestMessage_NotTo(t *testing.T) {
	_, err := msg.Text("你的快递已到，请携带工卡前往邮件中心领取。").Send()

	assertEqual(t, err.Error(), "toUser, toParty, toTag cannot be empty at the same time")
}

func TestMessage_Text(t *testing.T) {
	text := msg.Text().Content()

	resp, err := text.SetSafe(1).SetEnableIdTrans(0).Send()
	if err != nil {
		t.Error(err.Error())
	}

	if !equal(resp.Errcode, 0) {
		t.Error(resp.Errmsg)
	}

	assertEqual(t, text.safe, 1)
	assertEqual(t, text.enableIdTrans, 0)
}

func TestMessage_Image(t *testing.T) {
	image := msg.Image("3g2xoY5c_JUrK7-kQHW-__FXmsehOQyKX3116lrgx0jA")

	resp, err := image.Safe(1).Send()
	if err != nil {
		t.Error(err.Error())
	}

	if !equal(resp.Errcode, 0) {
		t.Error(resp.Errmsg)
	}

	assertEqual(t, image.safe, 1)
}

func TestMessage_Voice(t *testing.T) {
	voice := msg.Voice("37ucz7oOvPVWwrJBtMrj4xfPdrSuhlUsnxvBUnDixSYc6YvNAnVwRLVvN_KeWX1X8")

	resp, err := voice.Send()
	if err != nil {
		t.Error(err.Error())
	}

	if !equal(resp.Errcode, 0) {
		t.Error(resp.Errmsg)
	}
}

func TestMessage_Video(t *testing.T) {
	video := msg.Video("39tHzT7JokszltSEojcgjbJYO7tQqDOBlCL_wnj72kJXT3t0spBqt6_rlJNCHiyBM", "MP4", "mp4")

	resp, err := video.Safe(1).Send()
	if err != nil {
		t.Error(err.Error())
	}

	if !equal(resp.Errcode, 0) {
		t.Error(resp.Errmsg)
	}

	assertEqual(t, video.safe, 1)
}

func TestMessage_File(t *testing.T) {
	file := msg.File("3HILVuh11h1SDGTTMQ61p1rGcle0ZbNqBykIlXh37bLE")

	resp, err := file.Safe(1).Send()
	if err != nil {
		t.Error(err.Error())
	}

	if !equal(resp.Errcode, 0) {
		t.Error(resp.Errmsg)
	}

	assertEqual(t, file.safe, 1)
}

func TestMessage_Textcard(t *testing.T) {
	textcard := msg.Textcard("领奖通知", "<div class=\"gray\">2016年9月26日</div> <div class=\"normal\">恭喜你抽中iPhone 7一台，领奖码：xxxx</div><div class=\"highlight\">请于2016年10月10日前联系行政同事领取</div>", "http://work.weixin.qq.com")

	resp, err := textcard.Btntxt("详情").EnableIdTrans(0).Send()
	if err != nil {
		t.Error(err.Error())
	}

	if !equal(resp.Errcode, 0) {
		t.Error(resp.Errmsg)
	}

	assertEqual(t, textcard.btntxt, "详情")
	assertEqual(t, textcard.enableIdTrans, 0)
}

func TestMessage_Markdown(t *testing.T) {
	markdown := msg.Markdown().Content("您的会议室已经预定，稍后会同步到`邮箱` \n>**事项详情** \n>事　项：<font color=\"info\">开会</font> \n>组织者：@miglioguan \n>参与者：@miglioguan、@kunliu、@jamdeezhou、@kanexiong、@kisonwang \n> \n>会议室：<font color=\"info\">广州TIT 1楼 301</font> \n>日　期：<font color=\"warning\">2018年5月18日</font> \n>时　间：<font color=\"comment\">上午9:00-11:00</font> \n> \n>请准时参加会议。 \n> \n>如需修改会议信息，请点击：[修改会议信息](https://work.weixin.qq.com)").

	resp, err := markdown.Send()
	if err != nil {
		t.Error(err.Error())
	}

	if !equal(resp.Errcode, 0) {
		t.Error(resp.Errmsg)
	}
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
