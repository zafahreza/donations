package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Donations struct {
	Id          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	From        string             `json:"from,omitempty" bson:"from,omitempty"`
	Amount      int                `json:"amount,omitempty" bson:"amount,omitempty"`
	UserId      int                `json:"user_id,omitempty" bson:"user_id,omitempty"`
	Email       string             `json:"email,omitempty" bson:"email,omitempty"`
	Message     string             `json:"message,omitempty" bson:"message,omitempty"`
	IsAnonymous bool               `json:"is_anonymous,omitempty" bson:"is_anonymous,omitempty"`
	Payment     Payment            `json:"payment,omitempty" bson:"payment,omitempty"`
	CreatedAt   time.Time          `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt   time.Time          `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}

type Payment struct {
	TransactionStatus string `json:"transaction_status,omitempty" bson:"transaction_status,omitempty"`
	PaymentType       string `json:"payment_type,omitempty" bson:"payment_type,omitempty"`
	FraudStatus       string `json:"fraud_status,omitempty" bson:"fraud_status,omitempty"`
	TransactionTime   string `json:"transaction_time,omitempty" bson:"transaction_time,omitempty"`
	Currency          string `json:"currency,omitempty" bson:"currency,omitempty"`
	ExpireTime        string `json:"expire_time,omitempty" bson:"expire_time,omitempty"`
}
