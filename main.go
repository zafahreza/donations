package main

import (
	"context"
	"donations/alert"
	"donations/app"
	"donations/controller"
	"donations/model/domain"
	"donations/repository"
	"donations/service"
	"github.com/go-playground/validator/v10"
)

func main() {
	coll := app.NewSetupDB()
	validate := validator.New()
	kafkaReader := app.BrokerConsumer()
	kafkaWriter := app.BrokerProducer()
	alertWriter := app.BrokerProduserAlert()

	message := make(chan domain.User)
	usernameChan := make(chan string)
	newChanUsername := make(chan string)

	repo := repository.NewDonationRepository()
	alertService := alert.NewAlertService(alertWriter)
	paymentService := service.NewDonationPaymentService(alertService)
	services := service.NewDonationService(repo, coll, validate, paymentService)
	broker := service.NewDonationBrokerService(kafkaWriter, kafkaReader, message, usernameChan, newChanUsername)
	controllers := controller.NewDonationController(services, broker)

	go broker.ReadFromUser(context.Background())
	go broker.WriteToUser(context.Background())

	router := app.NewRouter(controllers)

	router.Run(":3000")
}
