package helper

import (
	"fmt"
	"github.com/kapsteur/event.club/config"
	"github.com/kapsteur/event.club/model"
	"log"
	"net"
	"net/mail"
	"net/smtp"
)

func SendMail(booking model.Booking) error {
	conf := config.Conf()

	body := ""

	body += fmt.Sprintf("Name:%s\n", booking.Name)
	body += fmt.Sprintf("Phone:%s\n", booking.Phone)
	body += fmt.Sprintf("Mail:%s\n", booking.Email)
	body += fmt.Sprintf("Midi:%b\n", booking.Meal1)
	body += fmt.Sprintf("Soir:%b\n", booking.Meal2)
	body += fmt.Sprintf("Comment:%s\n", booking.Comment)

	from := mail.Address{"", booking.Email}
	to := mail.Address{"", conf.Email.Email}

	// Setup headers
	headers := make(map[string]string)
	headers["From"] = from.String()
	headers["To"] = to.String()
	headers["Subject"] = "[Resa] " + booking.Name

	// Setup messages
	message := ""
	for k, v := range headers {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + body

	// Connect to the SMTP Server
	servername := "mail.gandi.net:587"

	host, _, _ := net.SplitHostPort(servername)

	auth := smtp.PlainAuth("", conf.Email.User, conf.Email.Password, host)

	// Here is the key, you need to call tls.Dial instead of smtp.Dial
	// for smtp servers running on 465 that require an ssl connection
	// from the very beginning (no starttls)
	//conn, err := tls.Dial("tcp", servername, tlsconfig)
	conn, err := net.Dial("tcp", servername)
	if err != nil {
		log.Println(err)
		return err
	}

	c, err := smtp.NewClient(conn, host)
	if err != nil {
		log.Println(err)
		return err
	}

	// Auth
	if err = c.Auth(auth); err != nil {
		log.Println(err)
		return err
	}

	// To && From
	if err = c.Mail(from.Address); err != nil {
		log.Println(err)
		return err
	}

	if err = c.Rcpt(to.Address); err != nil {
		log.Println(err)
		return err
	}

	// Data
	w, err := c.Data()
	if err != nil {
		log.Println(err)
		return err
	}

	_, err = w.Write([]byte(message))
	if err != nil {
		log.Println(err)
		return err
	}

	err = w.Close()
	if err != nil {
		log.Println(err)
		return err
	}

	c.Quit()
	return nil
}
