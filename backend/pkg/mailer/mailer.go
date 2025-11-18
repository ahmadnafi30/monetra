package mailer

import (
	"net/smtp"
	"os"
)

type Mailer interface {
	Send(to, subject, body string) error
}

type SMTPMailer struct{}

func NewSMTPMailer() *SMTPMailer {
	return &SMTPMailer{}
}

func (m *SMTPMailer) Send(to, subject, body string) error {
	from := os.Getenv("SMTP_FROM")
	username := os.Getenv("SMTP_USERNAME")
	password := os.Getenv("SMTP_PASSWORD")
	smtpHost := os.Getenv("SMTP_HOST")
	smtpPort := os.Getenv("SMTP_PORT")

	msg := []byte(
		"From: " + from + "\r\n" +
			"To: " + to + "\r\n" +
			"Subject: " + subject + "\r\n" +
			"MIME-Version: 1.0\r\n" +
			"Content-Type: text/html; charset=UTF-8\r\n" +
			"\r\n" + body,
	)

	auth := smtp.PlainAuth("", username, password, smtpHost)

	return smtp.SendMail(
		smtpHost+":"+smtpPort,
		auth,
		username,
		[]string{to},
		msg,
	)
}
