package mail

import (
	"fmt"
	"os"

	"github.com/go-mail/mail"
	"github.com/zaheerbabarkhan/todo-api-gogin/types"
)

func sendMail(mailData types.SendMailReq) error {
	newMail := mail.NewMessage()

	newMail.SetHeader("From", mailData.From)
	newMail.SetHeader("To", mailData.To)
	newMail.SetHeader("Subject", mailData.Subject)
	newMail.SetBody("text/html", mailData.Body)

	d := mail.NewDialer(os.Getenv("SMTP_HOST"), int(587), "sbabar950@gmail.com", os.Getenv("SMTP_PASSWORD"))

	if err := d.DialAndSend(newMail); err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func SendConfirmationEmail(to string, token string) error {
	mailData := types.SendMailReq{
		To:      to,
		From:    "bk885035@gmail.com",
		Subject: "Confirm Email",
		Body:    "Confirm Email",
	}
	err := sendMail(mailData)
	if err != nil {
		return err
	}
	return nil
}
