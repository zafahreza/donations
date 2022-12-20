package service

import (
	"context"
	"donations/helper"
	"donations/model/domain"
	"encoding/json"
	"fmt"
	"github.com/segmentio/kafka-go"
	"time"
)

type DonationBrokerImpl struct {
	Writer          *kafka.Writer
	Reader          *kafka.Reader
	Messages        chan domain.User
	UsernameChan    chan string
	NewChanUsername chan string
}

func NewDonationBrokerService(writer *kafka.Writer, reader *kafka.Reader, messages chan domain.User, usernameChain chan string, newChanUsername chan string) DonationBroker {
	return &DonationBrokerImpl{Writer: writer, Reader: reader, Messages: messages, UsernameChan: usernameChain, NewChanUsername: newChanUsername}
}

func (broker *DonationBrokerImpl) GetUserFromUser(ctx context.Context, username string) int {
	//fmt.Println("broker check point", time.Now())
	//
	//err := broker.Writer.WriteMessages(ctx,
	//	kafka.Message{
	//		Key:   []byte(username),
	//		Value: []byte(username),
	//	},
	//)
	//helper.PanicIfError(err)
	//
	//err = broker.Writer.Close()
	//helper.PanicIfError(err)
	//
	//fmt.Println("data sent to kafka", time.Now())
	broker.UsernameChan <- username

	user := <-broker.Messages

	return user.Id
}

func (broker *DonationBrokerImpl) WriteToUser(ctx context.Context) {

	for {
		username := <-broker.UsernameChan

		fmt.Println("broker check point", time.Now())
		err := broker.Writer.WriteMessages(ctx,
			kafka.Message{
				Key:   []byte(username),
				Value: []byte(username),
			},
		)
		helper.PanicIfError(err)

		//err = broker.Writer.Close()
		//helper.PanicIfError(err)

		fmt.Println("data sent to kafka", time.Now())
		//broker.NewChanUsername <- username
	}

}

func (broker *DonationBrokerImpl) ReadFromUser(ctx context.Context) {

	//defer broker.Reader.Close()
	for {
		err := broker.Reader.SetOffset(kafka.LastOffset)
		helper.PanicIfError(err)
		message, err := broker.Reader.ReadMessage(ctx)
		helper.PanicIfError(err)

		//key := <-broker.NewChanUsername
		//if string(message.Key) != key {
		//	fmt.Println("key salah")
		//	continue
		//}

		var newUser domain.User
		err = json.Unmarshal(message.Value, &newUser)
		helper.PanicIfError(err)

		broker.Messages <- newUser
	}
}
