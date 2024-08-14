package main

import (
	"emailSender/internal/domain/campaign"
	"emailSender/internal/infrastructure/database"
	"emailSender/internal/infrastructure/mail"
	"fmt"
	"log"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("error loading .env file")
	}

	db := database.NewDB()
	repository := database.CampaignRepository{Database: db}
	campaignService := campaign.ServiceImpl{
		Repository: &repository,
		SendEmail:  mail.SendEmail,
	}
	campaigns, err := repository.GetCampaignsToBeSent()

	if err != nil {
		fmt.Println(err.Error())
	}

	for {
		for _, campaign := range campaigns {
			campaignService.SendPendingEmail(&campaign)
		}

		time.Sleep(10 * time.Second)
	}
}
