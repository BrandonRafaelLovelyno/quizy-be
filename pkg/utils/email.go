package utils

import (
	"fmt"
	"quizy-be/internal/config"

	"gopkg.in/gomail.v2"
)

func SendHTMLEmail(config config.EmailConfig, to []string, subject, htmlBody string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", config.From)
	m.SetHeader("To", to...)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", htmlBody)

	d := gomail.NewDialer(config.Host, config.Port, config.Username, config.Password)

	if err := d.DialAndSend(m); err != nil {
		return fmt.Errorf("failed to send HTML email: %w", err)
	}
	return nil
}
