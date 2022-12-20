package helper

import (
	"donations/model/client"
)

func IsAnonymouse(donation client.DonationCreateRequest) string {
	isAnonymous := false
	isAnonymous = donation.IsAnonymous

	if isAnonymous == false {
		return donation.From
	}
	return "Seseorang"
}
