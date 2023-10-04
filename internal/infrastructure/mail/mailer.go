package mail

import (
	"email-dispatch-gateway/internal/domain/campaign"
	"github.com/Shopify/gomail"
	"os"
	"strconv"
)

type Mailer struct{}

func NewMailer() *Mailer {
	return &Mailer{}
}

func (m *Mailer) SendMail(c *campaign.Campaign) error {
	host := os.Getenv("EMAIL_HOST")
	port, _ := strconv.Atoi(os.Getenv("EMAIL_PORT"))
	username := os.Getenv("EMAIL_USERNAME")
	password := os.Getenv("EMAIL_PASSWORD")

	d := gomail.NewDialer(host, port, username, password)

	var emails []string
	for _, contact := range c.Contacts {
		emails = append(emails, contact.Email)
	}

	msg := gomail.NewMessage()
	msg.SetHeader("From", username)
	msg.SetHeader("To", emails...)
	msg.SetHeader("Subject", c.Name)
	msg.SetBody("text/html", c.Content)

	return d.DialAndSend(msg)
}
