package client

import (
	"time"
)

type DonationResponse struct {
	Id        string          `json:"id"`
	From      string          `json:"from"`
	To        int             `json:"to"`
	Amount    int             `json:"amount"`
	Message   string          `json:"message"`
	Email     string          `json:"email"`
	Payment   PaymentResponse `json:"payment"`
	UpdatedAt time.Time       `json:"updated_at"`
}

type PaymentResponse struct {
	TransactionStatus string `json:"transaction_status"`
	PaymentType       string `json:"payment_type"`
	TransactionTime   string `json:"transaction_time"`
	Currency          string `json:"currency"`
}

type DonationCreateResponse struct {
	Id         string    `json:"id"`
	From       string    `json:"from"`
	To         int       `json:"to"`
	Amount     int       `json:"amount"`
	Message    string    `json:"message"`
	Email      string    `json:"email"`
	PaymentUrl string    `json:"payment_url"`
	UpdatedAt  time.Time `json:"updated_at"`
}
