package service

import (
	"context"
)

type DonationBroker interface {
	GetUserFromUser(ctx context.Context, username string) int
	ReadFromUser(ctx context.Context)
	WriteToUser(ctx context.Context)
}
