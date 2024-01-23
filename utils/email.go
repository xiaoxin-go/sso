package utils

import (
	"crypto/tls"
	"fmt"
	"gopkg.in/gomail.v2"
	"sso/conf"
)

func SendEmail(subject string, body string, filename string, to ...string) error {
	data := conf.Config.Email
	email := gomail.NewMessage()
	email.SetHeader("From", data.Sender)                        // 发件人
	email.SetHeader("To", to...)                                // 发送给多个用户
	email.SetHeader("Subject", fmt.Sprintf("SSO: %s", subject)) // 邮件主题
	email.SetBody("text/html", body)                            // 邮件正文

	// 添加附件
	if len(filename) > 0 {
		email.Attach(filename)
	}
	fmt.Println(data.Host, data.Port, data.Username)
	d := gomail.NewDialer(data.Host, data.Port, data.Username, data.Password)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	err := d.DialAndSend(email)
	return err
}
