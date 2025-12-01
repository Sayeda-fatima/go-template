package common

import (
	"os"
	"strconv"

	"github.com/wneessen/go-mail"
)

type EmailService interface {
	SendEmail(to string, subject string, body string) error
}
type emailService struct {
	Host      string
	Port      string
	Username  string
	Password  string
	FromEmail string
}

func NewEmailService() EmailService {
	return &emailService{
		Host: os.Getenv("MAIL_HOST"),
		Port: os.Getenv("MAIL_PORT"),
		Username: os.Getenv("MAIL_USERNAME"),
		Password: os.Getenv("MAIL_PASSWORD"),
		FromEmail: os.Getenv("MAIL_FROM_ADDRESS"),
	}
}

func (es *emailService) SendEmail(to string, subject string, body string) error{

	port, _ := strconv.Atoi(es.Port)
	message := mail.NewMsg()

	if err := message.From(es.FromEmail); err != nil{
		return err
	}

	if err := message.To(to); err != nil{
		return err
	}

	message.Subject(subject)
	message.SetBodyString(mail.TypeTextHTML, body)

	client, err := mail.NewClient(es.Host, mail.WithSMTPAuth(mail.SMTPAuthPlain), mail.WithUsername(es.Username), mail.WithPassword(es.Password), mail.WithPort(port))
	
	if err != nil{
		return err
	}
	if err := client.DialAndSend(message); err != nil{
		return err
	}
	return nil
}