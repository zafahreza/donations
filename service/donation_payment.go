package service

import (
	"donations/model/client"
	"donations/model/domain"
	"github.com/gin-gonic/gin"
)

type DonationPayment interface {
	GetPaymentUrl(donation domain.Donations) string
	CheckTransaction(ctx *gin.Context, donation domain.Donations, request client.DonationNotification) domain.Donations
}
