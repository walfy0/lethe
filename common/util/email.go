package util

import (
	"github.com/lethe/common"
	"gopkg.in/gomail.v2"
)

func SendEmail(mail, body string) error {
	m := gomail.NewMessage()
	m.SetHeader("From",  m.FormatAddress(common.Email, "lethe"))
	m.SetHeader("To", mail)
	m.SetBody("text/html", body)
	d := gomail.NewDialer(common.YandexHost, common.YandexPort, common.Email, common.EmailPassWord)
	err := d.DialAndSend(m)
	return err
}
