package helpers

import (
	"bytes"
	"fmt"
	"text/template"

	"github.com/spf13/viper"
	"gopkg.in/gomail.v2"
)

type EmailRequest struct {
	Subject string
	To      string
	Content string
}

type EmailDataRequest struct {
	Name_customer  string
	package_name   string
	Counselor_name string
	Date_schedule  string
	Total          string
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

func SendingEmailWithHTML(to string, subject string, template string, request EmailDataRequest) error {
	result, errResults := ParseTemplateEmail("utils/html_template/"+template, request)

	if errResults != nil {
		return errResults
	}

	Message := gomail.NewMessage()
	Message.SetHeader("From", viper.GetString("email.USER"))
	Message.SetHeader("To", to)
	Message.SetHeader("Subject", subject)
	Message.SetBody("text/html", result)

	d := gomail.NewDialer(viper.GetString("EMAIL.HOST"), viper.GetInt("EMAIL.PORT"), viper.GetString("EMAIL.USER"), viper.GetString("EMAIL.PASS"))

	if errSending := d.DialAndSend(Message); errSending != nil {
		return errSending
	}

	return nil
}

func ParseTemplateEmail(templateFileName string, data interface{}) (string, error) {

	t, errTemplate := template.ParseFiles(templateFileName)

	if errTemplate != nil {
		return "", errTemplate
	}

	buf := new(bytes.Buffer)

	if errExce := t.Execute(buf, data); errExce != nil {
		return "", errExce
	}

	return buf.String(), nil

}
