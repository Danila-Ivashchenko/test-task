package main

import (
	"encoding/json"
	"fmt"
	"go-kafka/internal/domain/model"
	"log"

	"github.com/IBM/sarama"
)

func main() {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true

	// Создание продюсера
	producer, err := sarama.NewSyncProducer([]string{"localhost:9092"}, config)
	if err != nil {
		log.Fatalf("Error creating producer: %v", err)
	}
	defer producer.Close()

	dto := &model.UserInsertRequest{
		RawUser: model.RawUser{Name: "Oleg", Surname: "Tasg"},
		ToHash:  true,
	}
	b, err := json.Marshal(dto)
	if err != nil {
		panic(err)
	}
	message := &sarama.ProducerMessage{
		Topic: "FIO",
		Value: sarama.StringEncoder(string(b)),
	}

	// Отправка сообщения в тему Kafka
	partition, offset, err := producer.SendMessage(message)
	if err != nil {
		log.Fatalf("Error sending message: %v", err)
	}

	fmt.Printf("Message sent to partition %d at offset %d\n", partition, offset)
}
