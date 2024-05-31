package mail

import (
	"emailSender/internal/domain/campaign"
	"fmt"
	"os"

	"gopkg.in/gomail.v2"
)

func SendEmail(campaign *campaign.Campaign) error {
	fmt.Println("Sending mail...")

	message := gomail.NewMessage()

	d := gomail.NewDialer(os.Getenv("EMAIL_HOST"), 587, os.Getenv("EMAIL_USER"), os.Getenv("EMAIL_PASSWORD"))

	var emails []string

	for _, contact := range campaign.Contacts {
		emails = append(emails, contact.Email)
	}

	message.SetHeader("From", os.Getenv("EMAIL_USER"))
	message.SetHeader("To", emails...)
	message.SetHeader("Subject", campaign.Name)
	message.SetBody("text/html", campaign.Content)

	return d.DialAndSend(message)
}
