package sutil

import "testing"

func TestSend(t *testing.T) {
	mail := NewMail("smtp.office365.com", 25, "**********@example.com", "发件人昵称", "password")
	// 收件
	toers := make(map[string]string)
	toers["*************@163.com"] = "收件人昵称1"
	// 抄送
	ccers := make(map[string]string)
	ccers["*************@qq.com"] = "收件人昵称2"
	// 密送
	bccers := make(map[string]string)
	bccers["************@163.com"] = "收件人昵称3"
	msg := &MailMessage{
		Subject: "测试邮件",
		Body:    "<h1>您好，测试邮件，请勿回复！</h1>",
		Toers:   toers,
		Ccers:   ccers,
		Bccers:  bccers,
	}
	err := mail.Send(msg)
	if err != nil {
		t.Fatal(err)
	}
}
