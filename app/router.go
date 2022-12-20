package app

import (
	"donations/controller"
	"github.com/gin-gonic/gin"
)

func NewRouter(controller *controller.DonationController) *gin.Engine {

	router := gin.Default()
	api := router.Group("/api")

	api.POST("/donations/:username", controller.CreateDonation)
	api.GET("/donations/:user_id", controller.FindByUserId)
	api.POST("/donations/notification", controller.GetDonationNotification)

	return router
}
