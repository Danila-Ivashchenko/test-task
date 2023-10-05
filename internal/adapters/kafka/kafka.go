package kafka

import (
	"context"
	"encoding/json"
	"fmt"
	"go-kafka/internal/domain/model"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/IBM/sarama"
	"golang.org/x/exp/slog"
)

type userAdder interface {
	AddUser(context.Context, *model.UserInsertRequest) error
}

type config interface {
	GetBrokerAddrs() []string
	GetTopicToConsume() string
	GetTopicToProduce() string
	GetPartitionToConsume() int32
	GetTimeLimit() time.Duration
}

type kafkaAdapter struct {
	brokerAddrs        []string
	topicToConsume     string
	topicToProduce     string
	partitionToConsume int32

	config    *sarama.Config
	client    sarama.Client
	producer  sarama.SyncProducer
	consumer  sarama.Consumer
	partition sarama.PartitionConsumer

	userService userAdder
	timeLimit   time.Duration
	logger      *slog.Logger
}

func New(cfg config, u userAdder, l *slog.Logger) (*kafkaAdapter, error) {
	k := &kafkaAdapter{
		brokerAddrs:        cfg.GetBrokerAddrs(),
		topicToConsume:     cfg.GetTopicToConsume(),
		topicToProduce:     cfg.GetTopicToProduce(),
		partitionToConsume: cfg.GetPartitionToConsume(),
		config:             sarama.NewConfig(),
		userService:        u,
		timeLimit:          cfg.GetTimeLimit(),
		logger:             l,
	}
	k.config.Consumer.Return.Errors = true
	k.config.Producer.Return.Successes = true
	producer, err := sarama.NewSyncProducer(cfg.GetBrokerAddrs(), k.config)
	if err != nil {
		return nil, err
	}
	k.producer = producer

	client, err := sarama.NewClient(k.brokerAddrs, k.config)
	if err != nil {
		return nil, err
	}
	k.client = client

	consumer, err := sarama.NewConsumerFromClient(client)
	if err != nil {
		return nil, err
	}
	k.consumer = consumer

	return k, nil
}

func (k kafkaAdapter) Produce(msg error) error {
	message := &sarama.ProducerMessage{
		Topic: k.topicToProduce,
		Value: sarama.StringEncoder(msg.Error()),
	}
	_, _, err := k.producer.SendMessage(message)
	return err
}

func (k kafkaAdapter) Consume() error {
	log := k.logger.With(slog.String("component", "consumer"))
	part, err := k.consumer.ConsumePartition(k.topicToConsume, k.partitionToConsume, sarama.OffsetNewest)
	k.partition = part
	if err != nil {
		log.Error(("fail to consume"), slog.StringValue(err.Error()))
		return err
	}
	defer k.partition.Close()

	workerLog := log.With(
		slog.String("topic", k.topicToConsume),
		slog.String("partition", fmt.Sprintf("%d", k.partitionToConsume)),
	)

	workerLog.Info("consumer is available")

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
	for {
		select {
		case msg := <-part.Messages():
			t := time.Now()
			dto := &model.UserInsertRequest{}
			err := json.Unmarshal(msg.Value, &dto)
			if err != nil {
				workerLog.Error("error to parse data", slog.StringValue(err.Error()))
				k.Produce(err)
				break
			}

			context, cancel := context.WithTimeout(context.Background(), k.timeLimit)
			defer cancel()
			err = k.userService.AddUser(context, dto)
			offset := part.HighWaterMarkOffset()
			if err != nil {
				workerLog.Error(
					"action done with error",
					slog.String("error", err.Error()),
					slog.Int64("offser", offset),
					slog.String("duration", time.Since(t).String()),
				)
				k.Produce(err)
			} else {
				workerLog.Info(
					"action done",
					slog.Int64("offser", offset),
					slog.String("duration", time.Since(t).String()),
				)
			}
		case err := <-part.Errors():
			workerLog.Error("consumer over with error", slog.StringValue(err.Error()))
			return err
		case <-signals:
			workerLog.Info("consumer over")
			return nil
		}

	}
}
