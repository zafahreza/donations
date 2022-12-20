package test

import (
	"context"
	"donations/app"
	"donations/model/client"
	"donations/repository"
	"donations/service"
	"fmt"
	"github.com/go-playground/validator/v10"
	"testing"
)

func TestCreate(t *testing.T) {
	coll := app.NewSetupDB()
	valid := validator.New()
	repo := repository.NewDonationRepository()
	payment := service.NewDonationPaymentService()
	service := service.NewDonationService(repo, coll, valid, payment)
	ctx := context.Background()

	donation := client.DonationCreateRequest{
		From:    "fahreza",
		To:      32,
		Amount:  30000,
		Message: "donate lagi ahh",
		Email:   "fahreza@email.com",
	}

	result := service.Create(ctx, donation)
	fmt.Println(result)
}

func TestServiceFindById(t *testing.T) {
	coll := app.NewSetupDB()
	valid := validator.New()
	repo := repository.NewDonationRepository()
	payment := service.NewDonationPaymentService()
	service := service.NewDonationService(repo, coll, valid, payment)
	ctx := context.Background()

	result := service.FindById(ctx, "6391b637b1b2b032769c31d8")

	fmt.Println(result)
}
func TestFindByUserId(t *testing.T) {
	coll := app.NewSetupDB()
	valid := validator.New()
	repo := repository.NewDonationRepository()
	payment := service.NewDonationPaymentService()
	service := service.NewDonationService(repo, coll, valid, payment)
	ctx := context.Background()

	results := service.FindByUserId(ctx, 32)

	for _, result := range results {
		fmt.Println(result)
	}
}

func TestFindByEmail(t *testing.T) {
	coll := app.NewSetupDB()
	valid := validator.New()
	repo := repository.NewDonationRepository()
	payment := service.NewDonationPaymentService()
	service := service.NewDonationService(repo, coll, valid, payment)
	ctx := context.Background()

	results := service.FindByEmail(ctx, "fahreza@email.com")

	for _, result := range results {
		fmt.Println(result)
	}
}
