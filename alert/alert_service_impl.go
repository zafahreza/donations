package alert

import (
	"donations/model/domain"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/segmentio/kafka-go"
	"log"
)

type SendAlertServiceImpl struct {
	Writer *kafka.Writer
}

func NewAlertService(writer *kafka.Writer) SendAlertService {
	return &SendAlertServiceImpl{Writer: writer}
}

func (alert *SendAlertServiceImpl) SendToAlert(ctx *gin.Context, donation domain.Donations) {
	//key := strconv.Itoa(donation.UserId)

	result, err := json.Marshal(donation)
	if err != nil {
		log.Fatalln(err)
	}
	err = alert.Writer.WriteMessages(ctx,
		kafka.Message{
			Key:   []byte("coba coba aja"),
			Value: result,
		},
	)
	if err != nil {
		log.Fatalln(err)
	}
}
