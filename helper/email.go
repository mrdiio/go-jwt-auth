package helper

import (
	"bytes"
	"crypto/tls"
	"html/template"
	"log"
	"os"
	"path/filepath"

	"github.com/k3a/html2text"
	"github.com/spf13/viper"
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

func SendEmail(email string, emailData *EmailData) {
	to := email
	from := viper.GetString("EMAIL_FROM")

	smtpHost := viper.GetString("SMTP_PORT")
	smtpPort := viper.GetInt("SMTP_PORT")
	smtpPass := viper.GetString("SMTP_PASSWORD")
	smtpUser := viper.GetString("SMTP_USERNAME")

	var body bytes.Buffer
	temp, err := ParseTemplateDir("././templates")
	if err != nil {
		log.Fatal("Error parsing template: ", err)
	}

	temp.ExecuteTemplate(&body, "verificationCode.html", &email)

	m := gomail.NewMessage()
	m.SetHeader("From", from)
	m.SetHeader("To", to)
	m.SetHeader("Subject", emailData.Subject)
	m.SetBody("text/html", body.String())
	m.AddAlternative("text/plain", html2text.HTML2Text(body.String()))

	d := gomail.NewDialer(smtpHost, smtpPort, smtpUser, smtpPass)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	if err := d.DialAndSend(m); err != nil {
		log.Fatal("Error sending email: ", err)
	}

}
