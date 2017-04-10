// mail包提供简单的邮件发送功能。
package mail

import (
	"gopkg.in/gomail.v2"
)

// 简单邮件发送函数。
func Send(msg string) (err error) {
	m := gomail.NewMessage()
	m.SetHeader("From", "user@139.com")
	m.SetHeader("To", "user1@139.com", "user2@139.com")
	m.SetHeader("Subject", "【注意】某节点发现域名劫持")
	m.SetBody("text/html", msg)
	d := gomail.NewDialer("smtp.139.com", 465, "user@139.com", "password")
	if err = d.DialAndSend(m); err != nil {
		return
	}
	return
}
