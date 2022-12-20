package client

type DonationNotification struct {
	OrderId           string `json:"order_id"`
	TransactionStatus string `json:"transaction_status"`
	PaymentType       string `json:"payment_type"`
	FraudStatus       string `json:"fraud_status"`
	TransactionTime   string `json:"transaction_time"`
	Currency          string `json:"currency"`
	ExpireTime        string `json:"expire_time"`
}
