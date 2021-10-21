package wxcom

import (
	"os"
	"reflect"
	"strconv"
	"testing"
)

var agentid, _ = strconv.Atoi(os.Getenv("AGENTID"))
var message = New(os.Getenv("CORPID"), os.Getenv("CORPSECRET"), agentid).Message().SetTouser([]string{os.Getenv("TEST_USER")})

func TestMessage_Text(t *testing.T) {
	resp, err := message.Text("你的快递已到，请携带工卡前往邮件中心领取。").Send()
	if err != nil {
		t.Fatal(err.Error())
	}

	if !assertEqual(resp.Errcode, 0) {
		t.Fatal(resp.Errmsg)
	}
}

func TestMessage_Image(t *testing.T) {
	resp, err := message.Image("3g2xoY5c_JUrK7-kQHW-__FXmsehOQyKX3116lrgx0jA").Send()
	if err != nil {
		t.Fatal(err.Error())
	}

	if !assertEqual(resp.Errcode, 0) {
		t.Fatal(resp.Errmsg)
	}
}

func TestMessage_Voice(t *testing.T) {
	resp, err := message.Voice("37ucz7oOvPVWwrJBtMrj4xfPdrSuhlUsnxvBUnDixSYc6YvNAnVwRLVvN_KeWX1X8").Send()
	if err != nil {
		t.Fatal(err.Error())
	}

	if !assertEqual(resp.Errcode, 0) {
		t.Fatal(resp.Errmsg)
	}
}

func TestMessage_Video(t *testing.T) {
	resp, err := message.Video("39tHzT7JokszltSEojcgjbJYO7tQqDOBlCL_wnj72kJXT3t0spBqt6_rlJNCHiyBM", "MP4", "mp4").Send()
	if err != nil {
		t.Fatal(err.Error())
	}

	if !assertEqual(resp.Errcode, 0) {
		t.Fatal(resp.Errmsg)
	}
}

func TestMessage_File(t *testing.T) {
	resp, err := message.File("3HILVuh11h1SDGTTMQ61p1rGcle0ZbNqBykIlXh37bLE").Send()
	if err != nil {
		t.Fatal(err.Error())
	}

	if !assertEqual(resp.Errcode, 0) {
		t.Fatal(resp.Errmsg)
	}
}

func TestMessage_Textcard(t *testing.T) {
	resp, err := message.Textcard("领奖通知", "<div class=\"gray\">2016年9月26日</div> <div class=\"normal\">恭喜你抽中iPhone 7一台，领奖码：xxxx</div><div class=\"highlight\">请于2016年10月10日前联系行政同事领取</div>", "http://work.weixin.qq.com").Send()
	if err != nil {
		t.Fatal(err.Error())
	}

	if !assertEqual(resp.Errcode, 0) {
		t.Fatal(resp.Errmsg)
	}
}

func TestMessage_Markdown(t *testing.T) {
	resp, err := message.Markdown("您的会议室已经预定，稍后会同步到`邮箱` \n>**事项详情** \n>事　项：<font color=\"info\">开会</font> \n>组织者：@miglioguan \n>参与者：@miglioguan、@kunliu、@jamdeezhou、@kanexiong、@kisonwang \n> \n>会议室：<font color=\"info\">广州TIT 1楼 301</font> \n>日　期：<font color=\"warning\">2018年5月18日</font> \n>时　间：<font color=\"comment\">上午9:00-11:00</font> \n> \n>请准时参加会议。 \n> \n>如需修改会议信息，请点击：[修改会议信息](https://work.weixin.qq.com)").Send()
	if err != nil {
		t.Fatal(err.Error())
	}

	if !assertEqual(resp.Errcode, 0) {
		t.Fatal(resp.Errmsg)
	}
}

func assertEqual(e, g interface{}) (r bool) {
	if !equal(e, g) {
		return
	}

	return true
}

func equal(expected, got interface{}) bool {
	return reflect.DeepEqual(expected, got)
}
