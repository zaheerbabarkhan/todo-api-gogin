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

	// auth := smtp.PlainAuth("", "sbabar950@gmail.com", os.Getenv("SMTP_PASSWORD"), os.Getenv("SMTP_HOST"))

	// to := []string{mailData.To}

	// msg := []byte("To: " + mailData.To +

	// 	"Subject: Why aren’t you using Mailtrap yet?\r\n" +

	// 	"\r\n" +

	// 	"Here’s the space for our great sales pitch\r\n")

	// err := smtp.SendMail(os.Getenv("SMTP_HOST")+":"+fmt.Sprint(587), auth, "sbabar950@gmail.com", to, msg)

	// if err != nil {
	// 	fmt.Println(err)
	// 	return err
	// }
	// return nil
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
