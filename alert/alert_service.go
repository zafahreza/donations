package alert

import (
	"donations/model/domain"
	"github.com/gin-gonic/gin"
)

type SendAlertService interface {
	SendToAlert(ctx *gin.Context, donation domain.Donations)
}
