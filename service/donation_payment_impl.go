package service

import (
	"donations/alert"
	"donations/helper"
	"donations/model/client"
	"donations/model/domain"
	"github.com/gin-gonic/gin"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
	"time"
)

type DonationPaymentImpl struct {
	Alert alert.SendAlertService
}

func NewDonationPaymentService(alert alert.SendAlertService) DonationPayment {
	return &DonationPaymentImpl{
		Alert: alert,
	}
}

func (payment *DonationPaymentImpl) GetPaymentUrl(donation domain.Donations) string {
	var midtransClient snap.Client

	midtransClient.New("SB-Mid-server-Fg6xbTvgh5i2n7MK_-B8nPhW", midtrans.Sandbox)

	detail := midtrans.TransactionDetails{
		OrderID:  donation.Id.Hex(),
		GrossAmt: int64(donation.Amount),
	}

	customer := &midtrans.CustomerDetails{
		FName: donation.From,
		Email: donation.Email,
	}

	request := &snap.Request{
		TransactionDetails: detail,
		CustomerDetail:     customer,
	}

	snapReq, _ := midtransClient.CreateTransaction(request)

	return snapReq.RedirectURL

}

func (payment *DonationPaymentImpl) CheckTransaction(ctx *gin.Context, donation domain.Donations, request client.DonationNotification) domain.Donations {
	if request.TransactionStatus == "pending" {
		donation.Payment.PaymentType = request.PaymentType
		donation.Payment.TransactionStatus = helper.StatusPending()
		donation.Payment.FraudStatus = request.FraudStatus
		donation.Payment.TransactionTime = request.TransactionTime
		donation.Payment.ExpireTime = request.ExpireTime
		donation.Payment.Currency = request.Currency

		donation.UpdatedAt = time.Now()

		return donation
	}

	if request.TransactionStatus == "capture" || request.TransactionStatus == "settlement" {
		donation.Payment.PaymentType = request.PaymentType
		donation.Payment.TransactionStatus = helper.StatusPaid()
		donation.Payment.FraudStatus = request.FraudStatus
		donation.Payment.TransactionTime = request.TransactionTime
		donation.Payment.ExpireTime = request.ExpireTime
		donation.Payment.Currency = request.Currency

		donation.UpdatedAt = time.Now()

		payment.Alert.SendToAlert(ctx, donation)

		return donation
	}

	if request.FraudStatus == "deny" {
		donation.Payment.PaymentType = request.PaymentType
		donation.Payment.TransactionStatus = helper.StatusCancelled()
		donation.Payment.FraudStatus = request.FraudStatus
		donation.Payment.TransactionTime = request.TransactionTime
		donation.Payment.ExpireTime = request.ExpireTime
		donation.Payment.Currency = request.Currency

		donation.UpdatedAt = time.Now()

		return donation
	}

	if request.TransactionStatus == "deny" || request.TransactionStatus == "cancel" || request.TransactionStatus == "expire" {
		donation.Payment.PaymentType = request.PaymentType
		donation.Payment.TransactionStatus = helper.StatusCancelled()
		donation.Payment.FraudStatus = request.FraudStatus
		donation.Payment.TransactionTime = request.TransactionTime
		donation.Payment.ExpireTime = request.ExpireTime
		donation.Payment.Currency = request.Currency

		donation.UpdatedAt = time.Now()

		return donation
	}

	donation.Payment.PaymentType = request.PaymentType
	donation.Payment.TransactionStatus = helper.StatusPending()
	donation.Payment.FraudStatus = request.FraudStatus
	donation.Payment.TransactionTime = request.TransactionTime
	donation.Payment.ExpireTime = request.ExpireTime
	donation.Payment.Currency = request.Currency

	donation.UpdatedAt = time.Now()

	return donation
}
