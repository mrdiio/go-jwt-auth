package helper

import (
	"bytes"
	"crypto/tls"
	"html/template"
	"log"
	"os"
	"path/filepath"
	"strconv"

	"github.com/mrdiio/go-jwt-auth/models"
	"gopkg.in/gomail.v2"
)

type EmailData struct {
	URL       string
	FirstName string
	Subject   string
}

func ParseTemplateDir(dir string) (*template.Template, error) {
	var paths []string
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			paths = append(paths, path)
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return template.ParseFiles(paths...)
}

func SendEmail(user *models.User, email *EmailData) {
	from := os.Getenv("EMAIL_FROM")
	smtpHost := os.Getenv("SMTP_HOST")
	smtpPort := os.Getenv("SMTP_PORT")
	to := user.Email
	smtpPass := os.Getenv("SMTP_PASSWORD")
	smtpUser := os.Getenv("SMTP_USERNAME")

	port, err := strconv.Atoi(smtpPort)
	if err != nil {
		log.Fatal("Error parsing port: ", err)
	}

	var body bytes.Buffer
	temp, err := ParseTemplateDir("././templates")
	if err != nil {
		log.Fatal("Error parsing template: ", err)
	}

	temp.ExecuteTemplate(&body, "verificationCode.html", &email)

	m := gomail.NewMessage()
	m.SetHeader("From", from)
	m.SetHeader("To", *to)
	m.SetHeader("Subject", email.Subject)
	m.SetBody("text/html", body.String())
	// m.AddAlternative("text/plain", html2text.HTML2Text(body.String()))

	d := gomail.NewDialer(smtpHost, port, smtpUser, smtpPass)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	if err := d.DialAndSend(m); err != nil {
		log.Fatal("Error sending email: ", err)
	}

}
