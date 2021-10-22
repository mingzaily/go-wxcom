package wxcom

import (
	"fmt"
	"os"
	"testing"
	"time"
)

var testuser = os.Getenv("TEST_USER")

var message = wx.Message().SetTouser([]string{testuser})

func TestMessage(t *testing.T) {
	wx.cache.Set("access_token_"+fmt.Sprintf("%d", wx.agentid), "test_token", time.Minute)

	tempMessage := message.SetDuplicateCheck(1, 30)

	assertEqual(t, tempMessage.enableDuplicateCheck, 1)
	assertEqual(t, tempMessage.duplicateCheckInterval, 30)

	resp, err := tempMessage.Text("你的快递已到，请携带工卡前往邮件中心领取。").Send()
	if err != nil {
		t.Error(err.Error())
	}

	if !equal(resp.Errcode, 0) {
		t.Error(resp.Errmsg)
	}
}

func TestMessage_NotTo(t *testing.T) {
	tempMessage := wx.Message().SetTouser([]string{}).SetTotag([]string{}).SetToparty([]string{})

	_, err := tempMessage.Text("你的快递已到，请携带工卡前往邮件中心领取。").Send()

	assertEqual(t, err.Error(), "touser, toparty, totag cannot be empty at the same time")
}

func TestMessage_Text(t *testing.T) {
	text := message.Text("你的快递已到，请携带工卡前往邮件中心领取。")

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
	image := message.Image("3g2xoY5c_JUrK7-kQHW-__FXmsehOQyKX3116lrgx0jA")

	resp, err := image.SetSafe(1).Send()
	if err != nil {
		t.Error(err.Error())
	}

	if !equal(resp.Errcode, 0) {
		t.Error(resp.Errmsg)
	}

	assertEqual(t, image.safe, 1)
}

func TestMessage_Voice(t *testing.T) {
	voice := message.Voice("37ucz7oOvPVWwrJBtMrj4xfPdrSuhlUsnxvBUnDixSYc6YvNAnVwRLVvN_KeWX1X8")

	resp, err := voice.Send()
	if err != nil {
		t.Error(err.Error())
	}

	if !equal(resp.Errcode, 0) {
		t.Error(resp.Errmsg)
	}
}

func TestMessage_Video(t *testing.T) {
	video := message.Video("39tHzT7JokszltSEojcgjbJYO7tQqDOBlCL_wnj72kJXT3t0spBqt6_rlJNCHiyBM", "MP4", "mp4")

	resp, err := video.SetSafe(1).Send()
	if err != nil {
		t.Error(err.Error())
	}

	if !equal(resp.Errcode, 0) {
		t.Error(resp.Errmsg)
	}

	assertEqual(t, video.safe, 1)
}

func TestMessage_File(t *testing.T) {
	file := message.File("3HILVuh11h1SDGTTMQ61p1rGcle0ZbNqBykIlXh37bLE")

	resp, err := file.SetSafe(1).Send()
	if err != nil {
		t.Error(err.Error())
	}

	if !equal(resp.Errcode, 0) {
		t.Error(resp.Errmsg)
	}

	assertEqual(t, file.safe, 1)
}

func TestMessage_Textcard(t *testing.T) {
	textcard := message.Textcard("领奖通知", "<div class=\"gray\">2016年9月26日</div> <div class=\"normal\">恭喜你抽中iPhone 7一台，领奖码：xxxx</div><div class=\"highlight\">请于2016年10月10日前联系行政同事领取</div>", "http://work.weixin.qq.com")

	resp, err := textcard.SetBtntxt("详情").SetEnableIdTrans(0).Send()
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
	markdown := message.Markdown("您的会议室已经预定，稍后会同步到`邮箱` \n>**事项详情** \n>事　项：<font color=\"info\">开会</font> \n>组织者：@miglioguan \n>参与者：@miglioguan、@kunliu、@jamdeezhou、@kanexiong、@kisonwang \n> \n>会议室：<font color=\"info\">广州TIT 1楼 301</font> \n>日　期：<font color=\"warning\">2018年5月18日</font> \n>时　间：<font color=\"comment\">上午9:00-11:00</font> \n> \n>请准时参加会议。 \n> \n>如需修改会议信息，请点击：[修改会议信息](https://work.weixin.qq.com)")

	resp, err := markdown.Send()
	if err != nil {
		t.Error(err.Error())
	}

	if !equal(resp.Errcode, 0) {
		t.Error(resp.Errmsg)
	}
}
