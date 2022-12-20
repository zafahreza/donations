package service

import (
	"donations/model/client"
	"github.com/gin-gonic/gin"
)

type DonationsService interface {
	Create(ctx *gin.Context, request client.DonationCreateRequestServerSide) client.DonationCreateResponse
	FindById(ctx *gin.Context, donationId string) client.DonationResponse
	FindByUserId(ctx *gin.Context, userId int) []client.DonationResponse
	FindByEmail(ctx *gin.Context, email string) []client.DonationResponse
	UpdateStatus(ctx *gin.Context, request client.DonationNotification)
}
