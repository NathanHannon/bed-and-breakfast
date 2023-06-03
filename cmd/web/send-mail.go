package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"time"

	mail "github.com/xhit/go-simple-mail/v2"
	"nathanhannon.dev/bed-and-breakfast/internal/models"
)

func listenForMail() {
	go func() {
		for {
			message := <-app.MailChannel
			sendMessage(message)
		}
	}()
}

func sendMessage(m models.MailData) {
	const mailTimeout = 10 * time.Second
	server := mail.NewSMTPClient()
	server.Host = "localhost"
	server.Port = 1025
	server.KeepAlive = false
	server.ConnectTimeout = mailTimeout
	server.SendTimeout = mailTimeout

	client, err := server.Connect()
	if err != nil {
		errorLog.Println(err)
	}

	email := mail.NewMSG()
	email.SetFrom(m.From).AddTo(m.To).SetSubject(m.Subject)
	if m.Template == "" {
		email.SetBody(mail.TextHTML, m.Content)
	} else {
		data, err := ioutil.ReadFile(fmt.Sprintf("./email-templates/%s", m.Template))
		if err != nil {
			app.ErrorLog.Println(err)
		}
		mailTemplate := string(data)
		messageToSend := strings.Replace(mailTemplate, "[%body%]", m.Content, 1)
		email.SetBody(mail.TextHTML, messageToSend)
	}
	err = email.Send(client)
	if err != nil {
		errorLog.Println(err)
	} else {
		infoLog.Println("Email sent!")
	}
}
