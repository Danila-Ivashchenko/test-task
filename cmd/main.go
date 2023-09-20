package main

import (
	"fmt"
	"log"

	"github.com/IBM/sarama"
)

func main() {
	producer, err := sarama.NewSyncProducer([]string{"localhost:9092"}, nil)
	if err != nil {
		fmt.Println("error in connection")
		log.Fatal(err)
	}
	defer func() {
		if err := producer.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	topic := "test"
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder("SosI)!"),
	}

	partition, offset, err := producer.SendMessage(msg)

	if err != nil {
		fmt.Println("error in sending")
		log.Fatal(err)
	}

	fmt.Printf("Message sent to partition %d at offset %d\n", partition, offset)
}

// func main() {
// 	// Настройка потребителя Kafka
// 	config := sarama.NewConfig()
// 	config.Consumer.Return.Errors = true

// 	// Указание адресов брокеров Kafka
// 	brokers := []string{"localhost:9092"}
// 	consumer, err := sarama.NewConsumer(brokers, config)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer func() {
// 		if err := consumer.Close(); err != nil {
// 			panic(err)
// 		}
// 	}()
// 	fmt.Println(1)
// 	// Подписка на топики Kafka
// 	topics := []string{"test"}
// 	partitionConsumer, err := consumer.ConsumePartition(topics[0], 0, sarama.OffsetNewest)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer func() {
// 		if err := partitionConsumer.Close(); err != nil {
// 			panic(err)
// 		}
// 	}()
// 	fmt.Println(2)
// 	// Обработка сообщений из Kafka
// 	signals := make(chan os.Signal, 1)
// 	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
// 	fmt.Println(3)
// 	for {
// 		select {
// 		case msg := <-partitionConsumer.Messages():
// 			fmt.Printf("Received message from partition %d: %s\n", msg.Partition, string(msg.Value))
// 		case err := <-partitionConsumer.Errors():
// 			fmt.Printf("Error: %s\n", err.Error())
// 		case <-signals:
// 			return
// 		}
// 	}
// }
