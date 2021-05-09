package main

import (
	"time"

	"github.com/nathanhannon/bed-and-breakfast/internal/models"
	mail "github.com/xhit/go-simple-mail/v2"
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
	email.SetBody(mail.TextHTML, m.Content)
	err = email.Send(client)
	if err != nil {
		errorLog.Println(err)
	} else {
		infoLog.Println("Email sent!")
	}
}
