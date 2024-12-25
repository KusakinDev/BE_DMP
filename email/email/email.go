package email

import (
	emailconfig "back/email/emailConfig"

	"gopkg.in/gomail.v2"
)

func SendEmail(adressTo string, goodName string, content string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", emailconfig.Email) // Адрес отправителя
	m.SetHeader("To", adressTo)            // Адрес получателя
	m.SetHeader("Subject", "ДОСТАВКА ТОВАРА С САЙТА Alexander's digital marketplace!")

	message := "Ключ от " + goodName + ": " + content
	m.SetBody("text/plain", message)

	d := gomail.NewDialer(emailconfig.Host, emailconfig.Port, emailconfig.Email, emailconfig.Password)

	err := d.DialAndSend(m)

	return err
}
