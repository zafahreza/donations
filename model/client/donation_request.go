package client

type DonationCreateRequest struct {
	From        string `validate:"required" json:"from"`
	Amount      int    `validate:"required" json:"amount"`
	Email       string `validate:"required,email" json:"email"`
	Message     string `json:"message"`
	IsAnonymous bool   `json:"is_anonymous"`
}

type DonationCreateRequestServerSide struct {
	From        string `validate:"required" json:"from"`
	To          int    `validate:"required" json:"to"`
	Amount      int    `validate:"required" json:"amount"`
	Email       string `validate:"required,email" json:"email"`
	Message     string `json:"message"`
	IsAnonymous bool   `json:"is_anonymous"`
}
