package test

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"log"
	"testing"
)

func TestKafkaReader(t *testing.T) {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{"localhost:9092"},
		Topic:     "contoh",
		Partition: 0,    // 10KB
		MaxBytes:  10e6, // 10MB
	})

	for {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			break
		}
		fmt.Printf("message at offset %d: %s = %s\n", m.Offset, string(m.Key), string(m.Value))
	}

	if err := r.Close(); err != nil {
		log.Fatal("failed to close reader:", err)
	}
}

func TestKafkaWriter(t *testing.T) {

	//topic := "contoh"
	//partition := 0
	//
	//conn, err := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", topic, partition)
	//if err != nil {
	//	log.Fatal("failed to dial leader:", err)
	//}
	// make a writer that produces to topic-A, using the least-bytes distribution
	w := &kafka.Writer{
		Addr:  kafka.TCP("localhost:9092"),
		Topic: "contoh",
	}

	err := w.WriteMessages(context.Background(),
		kafka.Message{
			Key:   []byte("Key-A"),
			Value: []byte("test"),
		},
		kafka.Message{
			Key:   []byte("Key-B"),
			Value: []byte("lagi"),
		},
		kafka.Message{
			Key:   []byte("Key-C"),
			Value: []byte("coba"),
		},
	)
	if err != nil {
		log.Fatal("failed to write messages:", err)
	}

	if err := w.Close(); err != nil {
		log.Fatal("failed to close writer:", err)
	}

}
