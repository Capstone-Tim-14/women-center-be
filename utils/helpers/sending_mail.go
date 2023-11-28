package helpers

import (
	"fmt"

	"github.com/spf13/viper"
	"gopkg.in/gomail.v2"
)

type EmailRequest struct {
	Subject string
	To      string
	Content string
}

func SendingEmail(request EmailRequest) error {

	mail := gomail.NewMessage()

	mail.SetHeader("From", viper.GetString("EMAIL.USER"))
	mail.SetHeader("To", request.To)
	mail.SetHeader("Subject", request.Subject)
	mail.SetBody("text/html", request.Content)

	d := gomail.NewDialer(viper.GetString("EMAIL.HOST"), viper.GetInt("EMAIL.PORT"), viper.GetString("EMAIL.USER"), viper.GetString("EMAIL.PASS"))

	if errDialeg := d.DialAndSend(mail); errDialeg != nil {
		fmt.Println(errDialeg)
		return errDialeg
	}

	return nil

}
