package helper

import (
	"donations/model/client"
	"donations/model/domain"
)

func ToDonationResponse(donation domain.Donations) client.DonationResponse {
	if donation.IsAnonymous == true {
		donation.From = "Seseorang"
	}

	payment := client.PaymentResponse{
		TransactionStatus: donation.Payment.TransactionStatus,
		PaymentType:       donation.Payment.PaymentType,
		TransactionTime:   donation.Payment.TransactionTime,
		Currency:          donation.Payment.Currency,
	}

	response := client.DonationResponse{
		Id:        donation.Id.Hex(),
		From:      donation.From,
		To:        donation.UserId,
		Amount:    donation.Amount,
		Message:   donation.Message,
		Email:     donation.Email,
		Payment:   payment,
		UpdatedAt: donation.UpdatedAt,
	}

	return response
}

func ToDonationsResponses(donations []domain.Donations) []client.DonationResponse {
	var responses []client.DonationResponse

	for _, donation := range donations {
		response := ToDonationResponse(donation)
		responses = append(responses, response)
	}

	return responses
}

func ToDonationCreateResponse(donation domain.Donations, paymentUrl string) client.DonationCreateResponse {
	response := client.DonationCreateResponse{
		Id:         donation.Id.Hex(),
		From:       donation.From,
		To:         donation.UserId,
		Amount:     donation.Amount,
		Message:    donation.Message,
		Email:      donation.Email,
		PaymentUrl: paymentUrl,
		UpdatedAt:  donation.UpdatedAt,
	}

	return response
}

func ToDonationCreateRequestServerSide(request client.DonationCreateRequest, userId int) client.DonationCreateRequestServerSide {
	result := client.DonationCreateRequestServerSide{
		From:        request.From,
		To:          userId,
		Amount:      request.Amount,
		Email:       request.Email,
		Message:     request.Message,
		IsAnonymous: request.IsAnonymous,
	}

	return result
}
