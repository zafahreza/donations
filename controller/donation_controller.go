package controller

import (
	"donations/exception"
	"donations/helper"
	"donations/model/client"
	"donations/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"net/http"
)

type DonationController struct {
	Service service.DonationsService
	Broker  service.DonationBroker
}

func NewDonationController(service service.DonationsService, broker service.DonationBroker) *DonationController {
	return &DonationController{Service: service, Broker: broker}
}

func (controller *DonationController) CreateDonation(c *gin.Context) {
	var request client.DonationCreateRequest
	var username client.CreateDonationUri

	err := c.ShouldBindUri(&username)
	if err != nil {
		exception.NewNotFoundError(c, errors.New("invalid url"))
		return
	}

	fmt.Println("check point")

	userId := controller.Broker.GetUserFromUser(c, username.Username)
	if userId == 0 {
		exception.NewNotFoundError(c, errors.New("invalid url"))
		return
	}

	err = c.ShouldBindJSON(&request)
	exception.Unprocessable(c, err)

	goodRequest := helper.ToDonationCreateRequestServerSide(request, userId)

	responseData := controller.Service.Create(c, goodRequest)
	ResponseBody := client.HttpResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   responseData,
	}
	c.JSON(http.StatusOK, ResponseBody)
}

func (controller *DonationController) FindByUserId(c *gin.Context) {
	var userId client.GetDonationUri

	err := c.ShouldBindUri(&userId)
	helper.PanicIfError(err)

	responseData := controller.Service.FindByUserId(c, userId.UserId)
	responseBody := client.HttpResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   responseData,
	}
	c.JSON(http.StatusOK, responseBody)
}

func (controller *DonationController) GetDonationNotification(c *gin.Context) {
	var notif client.DonationNotification

	err := c.ShouldBindJSON(&notif)
	helper.PanicIfError(err)

	controller.Service.UpdateStatus(c, notif)

	c.JSON(http.StatusOK, nil)
}
