package test

import (
	"cloud-disk/core/define"
	"crypto/tls"
	"github.com/jordan-wright/email"
	"net/smtp"
	"testing"
)

func TestSendEmail(t *testing.T) {
	e := email.NewEmail()
	e.From = "Get <me_rhythm@163.com>"
	e.To = []string{"1256157634@qq.com"}
	e.Subject = "发送验证码测试"
	e.HTML = []byte("您的验证码是：<b>123456</b>")
	err := e.SendWithTLS("smtp.163.com:465",
		smtp.PlainAuth("", "me_rhythm@163.com", define.MailPassword, "smtp.163.com"),
		&tls.Config{InsecureSkipVerify: true, ServerName: "smtp.163.com"})
	if err != nil {
		t.Fatal(err)
	}
}
