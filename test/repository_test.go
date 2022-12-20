package test

import (
	"context"
	"donations/app"
	"donations/model/domain"
	"donations/repository"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"testing"
	"time"
)

func TestInsert(t *testing.T) {
	donation := domain.Donations{
		From:        "yoolanda",
		Amount:      32000,
		UserId:      32,
		Email:       "yolanda@email.com",
		Message:     "mencoba donation lagi buat contoh data",
		IsAnonymous: true,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	coll := app.NewSetupDB()

	ctx := context.Background()

	repo := repository.NewDonationRepository()

	donationResult := repo.Insert(ctx, coll, donation)

	fmt.Println(donationResult)
}

func TestFindOne(t *testing.T) {
	coll := app.NewSetupDB()

	ctx := context.Background()

	repo := repository.NewDonationRepository()

	filter := bson.D{{"from", "yoolanda"}}

	result, err := repo.FindOne(ctx, coll, filter)
	if err != nil {
		fmt.Println("no document found")
		return
	}

	stringId := result.Id.Hex()

	fmt.Println(stringId)
	fmt.Println(result.From)
	fmt.Println(result.Amount)
	fmt.Println(result.UserId)
	fmt.Println(result.Email)
	fmt.Println(result.Message)
	fmt.Println(result.IsAnonymous)
	fmt.Println(result.Payment)
	fmt.Println(result.CreatedAt)
}

func TestFindById(t *testing.T) {
	coll := app.NewSetupDB()

	ctx := context.Background()

	repo := repository.NewDonationRepository()

	result, err := repo.FindById(ctx, coll, "6392ed1fc363d36e952d01c7")
	if err != nil {
		fmt.Println("no document found", err)
		return
	}

	stringId := result.Id.Hex()

	fmt.Println(stringId)
	fmt.Println(result.From)
	fmt.Println(result.Amount)
	fmt.Println(result.UserId)
	fmt.Println(result.Email)
	fmt.Println(result.Message)
	fmt.Println(result.IsAnonymous)
	fmt.Println(result.Payment)
	fmt.Println(result.CreatedAt)
}

func TestUpdateById(t *testing.T) {
	coll := app.NewSetupDB()

	ctx := context.Background()

	repo := repository.NewDonationRepository()

	resultFind, err := repo.FindById(ctx, coll, "6390bcf1e66deba844b15f77")
	if err != nil {
		fmt.Println("no document found")
		return
	}

	result, err := repo.UpdateById(ctx, coll, "63919f11dc7269c5861d8e07", domain.Donations{UpdatedAt: time.Now()})
	if err != nil {
		fmt.Println("error disini", err)
		return
	}

	stringId := resultFind.Id.Hex()

	fmt.Println(stringId)
	fmt.Println(resultFind.From)
	fmt.Println(resultFind.Amount)
	fmt.Println(resultFind.UserId)
	fmt.Println(resultFind.Email)
	fmt.Println(resultFind.Message)
	fmt.Println(resultFind.IsAnonymous)
	fmt.Println(result.Payment)
	fmt.Println(resultFind.CreatedAt)
	fmt.Println(result.UpdatedAt)
}

func TestFind(t *testing.T) {
	coll := app.NewSetupDB()

	ctx := context.Background()

	repo := repository.NewDonationRepository()

	filter := bson.M{"status": "Paid"}

	result, err := repo.Find(ctx, coll, filter)
	if err != nil {
		fmt.Println("no documment found")
		return
	}

	for _, donation := range result {
		fmt.Println(donation)
	}
}
