package service

import (
	"donations/exception"
	"donations/helper"
	"donations/model/client"
	"donations/model/domain"
	"donations/repository"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"time"
)

type DonationServiceImpl struct {
	Repository repository.DonationsRepository
	Coll       *mongo.Collection
	Validate   *validator.Validate
	Payment    DonationPayment
}

func NewDonationService(repository repository.DonationsRepository, coll *mongo.Collection, validate *validator.Validate, payment DonationPayment) DonationsService {
	return &DonationServiceImpl{Repository: repository, Coll: coll, Validate: validate, Payment: payment}
}

func (service *DonationServiceImpl) Create(ctx *gin.Context, request client.DonationCreateRequestServerSide) client.DonationCreateResponse {
	err := service.Validate.Struct(request)
	exception.Unprocessable(ctx, err)

	coll := service.Coll

	payment := domain.Payment{
		TransactionStatus: helper.StatusPending(),
	}

	donation := domain.Donations{
		From:        request.From,
		UserId:      request.To,
		Amount:      request.Amount,
		Message:     request.Message,
		Email:       request.Email,
		IsAnonymous: request.IsAnonymous,
		Payment:     payment,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	newDonation := service.Repository.Insert(ctx, coll, donation)

	paymentUrl := service.Payment.GetPaymentUrl(newDonation)

	return helper.ToDonationCreateResponse(newDonation, paymentUrl)

}

func (service *DonationServiceImpl) FindById(ctx *gin.Context, donationId string) client.DonationResponse {
	coll := service.Coll

	donation, err := service.Repository.FindById(ctx, coll, donationId)
	helper.PanicIfError(err)

	return helper.ToDonationResponse(donation)
}

func (service *DonationServiceImpl) FindByUserId(ctx *gin.Context, userId int) []client.DonationResponse {
	coll := service.Coll
	filter := bson.M{"$and": bson.A{
		bson.D{{"user_id", userId}},
		bson.D{{"payment.transaction_status", "paid"}},
	}}

	donations, err := service.Repository.Find(ctx, coll, filter)
	if err != nil {
		exception.NewNotFoundError(ctx, errors.New("not found"))
	}

	return helper.ToDonationsResponses(donations)
}

func (service *DonationServiceImpl) FindByEmail(ctx *gin.Context, email string) []client.DonationResponse {
	coll := service.Coll
	filter := bson.M{"email": email}

	donations, err := service.Repository.Find(ctx, coll, filter)
	helper.PanicIfError(err)

	return helper.ToDonationsResponses(donations)
}

func (service *DonationServiceImpl) UpdateStatus(ctx *gin.Context, request client.DonationNotification) {
	coll := service.Coll

	donation, err := service.Repository.FindById(ctx, coll, request.OrderId)
	if err != nil {
		log.Fatalln(err)
	}

	result := service.Payment.CheckTransaction(ctx, donation, request)

	_, err = service.Repository.UpdateById(ctx, coll, request.OrderId, result)
	exception.InternalServerError(ctx, err)
}
