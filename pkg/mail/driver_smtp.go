package mail

import (
	"codeReleaseSystem/pkg/logger"
	"fmt"
	"net/smtp"

	emailPKG "github.com/jordan-wright/email"
)

// SMTP 实现 email.Driver interface
type SMTP struct{}

// Send 实现 email.Driver interface 的 Send 方法
func (s *SMTP) Send(email Email, config map[string]string) bool {
	e := emailPKG.NewEmail()

	e.From = fmt.Sprintf("%v <%v>", email.From.Name, email.From.Address)
	e.To = email.To
	e.Bcc = email.Bcc
	e.Cc = email.Cc
	e.Subject = email.Subject
	e.Text = email.Text
	e.HTML = email.Html

	logger.DebugJSON("Mail", "邮件发送配置", e)

	// 发送邮件
	err := e.Send(fmt.Sprintf("%v:%v", config["host"], config["port"]), smtp.PlainAuth("", config["username"], config["password"], config["host"]))
	if err != nil {
		logger.ErrorString("Mail", "邮件发送失败", err.Error())
		return false
	}

	logger.DebugString("Mail", "邮件发送成功", fmt.Sprintf("邮件发送成功，发件人：%v，收件人：%v", email.From.Address, email.To))
	return true
}
