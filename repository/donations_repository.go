package repository

import (
	"context"
	"donations/model/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type DonationsRepository interface {
	Insert(ctx context.Context, coll *mongo.Collection, donation domain.Donations) domain.Donations
	FindOne(ctx context.Context, coll *mongo.Collection, filter bson.D) (domain.Donations, error)
	Find(ctx context.Context, coll *mongo.Collection, filter bson.M) ([]domain.Donations, error)
	FindById(ctx context.Context, coll *mongo.Collection, filter string) (domain.Donations, error)
	UpdateById(ctx context.Context, coll *mongo.Collection, filter string, set domain.Donations) (domain.Donations, error)
}
