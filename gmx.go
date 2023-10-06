package main

import (
	"fmt"
	"net/smtp"
	"strings"
)

func main() {
	// Sender's email credentials
	from := "username@gmx.com"
	password := "<password>"

	// Recipient email address
	to := "example@hotmail.com"

	// SMTP server details for GMX
	smtpHost := "mail.gmx.com"
	smtpPort := 587

	// Message content
	subject := "Subject of your email"
	body := "This is the body of your email."

	// Compose the email message
	message := composeEmail(from, to, subject, body)

	// Authentication credentials
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Connect to the SMTP server
	err := smtp.SendMail(
		fmt.Sprintf("%s:%d", smtpHost, smtpPort),
		auth,
		from,
		[]string{to},
		[]byte(message),
	)
	if err != nil {
		fmt.Println("Error sending email:", err)
		return
	}

	fmt.Println("Email sent successfully!")
}

func composeEmail(from, to, subject, body string) string {
	headers := make(map[string]string)
	headers["From"] = from
	headers["To"] = to
	headers["Subject"] = subject
	headers["MIME-Version"] = "1.0"
	headers["Content-Type"] = "text/plain; charset=UTF-8"
	headers["Content-Transfer-Encoding"] = "base64"

	var headerLines []string
	for key, value := range headers {
		headerLines = append(headerLines, fmt.Sprintf("%s: %s", key, value))
	}

	return fmt.Sprintf("%s\r\n\r\n%s", strings.Join(headerLines, "\r\n"), body)
}
