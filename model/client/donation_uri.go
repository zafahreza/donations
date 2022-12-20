package client

type GetDonationUri struct {
	UserId int `uri:"user_id" binding:"required"`
}

type CreateDonationUri struct {
	Username string `uri:"username" binding:"required"`
}
