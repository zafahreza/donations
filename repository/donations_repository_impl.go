package repository

import (
	"context"
	"donations/helper"
	"donations/model/domain"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type DonationsRepositoryImpl struct {
}

func NewDonationRepository() DonationsRepository {
	return &DonationsRepositoryImpl{}
}

func (repository *DonationsRepositoryImpl) Insert(ctx context.Context, coll *mongo.Collection, donation domain.Donations) domain.Donations {
	result, err := coll.InsertOne(ctx, donation)
	if err != nil {
		log.Fatal(err)
	}

	donation.Id = result.InsertedID.(primitive.ObjectID)
	return donation
}

func (repository *DonationsRepositoryImpl) FindOne(ctx context.Context, coll *mongo.Collection, filter bson.D) (domain.Donations, error) {
	var donation domain.Donations

	err := coll.FindOne(ctx, filter).Decode(&donation)
	if err != nil {
		return donation, err
	}

	return donation, nil
}

func (repository *DonationsRepositoryImpl) FindById(ctx context.Context, coll *mongo.Collection, filter string) (domain.Donations, error) {
	id, err := primitive.ObjectIDFromHex(filter)
	helper.PanicIfError(err)

	var donation domain.Donations
	err = coll.FindOne(ctx, bson.D{{"_id", id}}).Decode(&donation)
	if err != nil {
		fmt.Println(id)
		fmt.Println(donation)
		return donation, err
	}

	return donation, nil

}

func (repository *DonationsRepositoryImpl) UpdateById(ctx context.Context, coll *mongo.Collection, filter string, set domain.Donations) (domain.Donations, error) {
	id, err := primitive.ObjectIDFromHex(filter)
	helper.PanicIfError(err)

	_, err = coll.UpdateByID(ctx, id, bson.M{"$set": set})
	if err != nil {
		return set, err
	}

	return set, nil

}

func (repository *DonationsRepositoryImpl) Find(ctx context.Context, coll *mongo.Collection, filter bson.M) ([]domain.Donations, error) {
	var donations []domain.Donations

	opts := options.Find().SetSort(bson.D{{"updated_at", -1}}).SetLimit(10).SetSkip(0)

	cursor, err := coll.Find(ctx, filter, opts)
	if err != nil {
		return donations, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var donation domain.Donations
		cursor.Decode(&donation)
		donations = append(donations, donation)
	}

	return donations, nil
}
