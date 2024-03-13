package utils

import (
	"log"
	"net/smtp"
	"os"
)

func SendEmail(to string, subject string, body string) error {
	authmail := os.Getenv("AUTH_MAIL")
	from := os.Getenv("MAIL_FROM")
	password := os.Getenv("MAIL_AUTH_PASS")
	smtp_server_addr := os.Getenv("SMTP_SERVER_ADDR")
	smtp_server := os.Getenv("SMTP_SERVER")
	// メールヘッダーの作成
	header := make(map[string]string)
	header["From"] = "CharisWorks本部" + " <" + from + ">"
	header["To"] = to
	header["Subject"] = subject

	// メール本文の作成
	message := ""
	for key, value := range header {
		message += key + ": " + value + "\r\n"
	}
	message += "\r\n" + body
	err := smtp.SendMail(smtp_server_addr,
		smtp.PlainAuth("本部", authmail, password, smtp_server),
		from, []string{to}, []byte(message))
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Email sent successfully!")
	}
	return nil
}
