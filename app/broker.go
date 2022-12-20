package app

import (
	"github.com/segmentio/kafka-go"
)

func BrokerProducer() *kafka.Writer {
	w := &kafka.Writer{
		Addr:  kafka.TCP("localhost:9092"),
		Topic: "username",
	}

	return w
}

func BrokerConsumer() *kafka.Reader {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:     []string{"localhost:9092"},
		Topic:       "userDonation",
		Partition:   0,
		MaxBytes:    10e6, // 10MB
		StartOffset: kafka.LastOffset,
	})

	return r
}

func BrokerProduserAlert() *kafka.Writer {
	w := &kafka.Writer{
		Addr:  kafka.TCP("localhost:9092"),
		Topic: "donationAlert",
	}

	return w
}
