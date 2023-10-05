package config

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type config struct {
	postgresUser       string
	postgresPass       string
	postgresHost       string
	postgresPort       string
	postgresDB         string
	postgresSSLMode    string
	env                string
	httpPort           string
	httpHost           string
	historyDir         string
	brokerHost         string
	brokerPort         string
	topicToConsume     string
	topicToProduce     string
	partitionToConsume int32
	timeLimit          time.Duration
	redisHost          string
	redisPort          string
	redisDb            int
	redisPassword      string
	redisTtl           time.Duration
}

func (cfg config) GetPostgresUser() string {
	return cfg.postgresUser
}

func (cfg config) GetPostgresPass() string {
	return cfg.postgresPass
}

func (cfg config) GetPostgresHost() string {
	return cfg.postgresHost
}

func (cfg config) GetPostgresPort() string {
	return cfg.postgresPort
}

func (cfg config) GetPostgresDB() string {
	return cfg.postgresDB
}

func (cfg config) GetPostgresSSLMode() string {
	return cfg.postgresSSLMode
}

func (cfg config) GetEnv() string {
	return cfg.env
}

func (cfg config) GetHttpPort() string {
	return cfg.httpPort
}

func (cfg config) GetHttpURL() string {
	return cfg.httpHost + ":" + cfg.httpPort
}
func (cfg config) GetHistoryDir() string {
	return cfg.historyDir
}

func (cfg config) GetBrokerAddrs() []string {
	return []string{
		cfg.brokerHost + ":" + cfg.brokerPort,
	}
}
func (cfg config) GetTopicToConsume() string {
	return cfg.topicToConsume
}
func (cfg config) GetTopicToProduce() string {
	return cfg.topicToProduce
}
func (cfg config) GetPartitionToConsume() int32 {
	return cfg.partitionToConsume
}
func (cfg config) GetTimeLimit() time.Duration {
	return cfg.timeLimit
}
func (cfg config) GetRedisTtl() time.Duration {
	return cfg.redisTtl
}
func (cfg config) GetRedisPort() string {
	return cfg.redisPort
}
func (cfg config) GetRedisHost() string {
	return cfg.redisHost
}
func (cfg config) GetRedisPassword() string {
	return cfg.redisPassword
}
func (cfg config) GetRedisDB() int {
	return cfg.redisDb
}

func LoadEnv(filenames ...string) error {
	const op = "pkg.config.LoadEnv"
	err := godotenv.Load(filenames...)
	if err != nil {
		return fmt.Errorf("%s: %s", op, err)
	}
	return nil
}

func GetConfig() *config {
	cfg := &config{
		postgresUser:    "",
		postgresPass:    "",
		postgresHost:    "localhost",
		postgresPort:    "27017",
		postgresDB:      "",
		env:             "local",
		postgresSSLMode: "disable",
		httpHost:        "localhost",
		historyDir:      "history",
		timeLimit:       time.Second * 10,
		redisHost:       "localhost",
		redisPort:       "6379",
		redisDb:         0,
		redisTtl:        time.Hour * 24,
	}

	user := os.Getenv("POSTGRES_USER")
	pass := os.Getenv("POSTGRES_PASSWORD")
	host := os.Getenv("POSTGRES_HOST")
	port := os.Getenv("POSTGRES_PORT")
	db := os.Getenv("POSTGRES_DB")
	ssl := os.Getenv("POSTGRES_SSL_MODE")
	env := os.Getenv("ENV")
	httpPort := os.Getenv("HTTP_PORT")
	httpHost := os.Getenv("HTTP_HOST")
	brokerHost := os.Getenv("BROKER_HOST")
	brokerPort := os.Getenv("BROKER_PORT")
	topicToConsume := os.Getenv("TOPIC_TO_CONSUME")
	topicToProduce := os.Getenv("TOPIC_TO_PRODUCE")
	partitionToConsume := os.Getenv("PARTITION_TO_CONSUME")
	timeLimit := os.Getenv("TIME_LIMIT")

	if env != "" {
		cfg.env = env
	}
	if httpPort != "" {
		cfg.httpPort = httpPort
	}
	if httpHost != "" {
		cfg.httpHost = httpHost
	}
	if user != "" {
		cfg.postgresUser = user
	}
	if pass != "" {
		cfg.postgresPass = pass
	}
	if host != "" {
		cfg.postgresHost = host
	}
	if port != "" {
		cfg.postgresPort = port
	}
	if db != "" {
		cfg.postgresDB = db
	}
	if ssl != "" {
		cfg.postgresSSLMode = ssl
	}
	if brokerHost != "" {
		cfg.brokerHost = brokerHost
	}
	if brokerPort != "" {
		cfg.brokerPort = brokerPort
	}
	if topicToConsume != "" {
		cfg.topicToConsume = topicToConsume
	}
	if topicToProduce != "" {
		cfg.topicToProduce = topicToProduce
	}
	if partitionToConsume != "" {
		val, err := strconv.Atoi(partitionToConsume)
		if err == nil {
			cfg.partitionToConsume = int32(val)
		}

	}
	if timeLimit != "" {
		val, err := strconv.Atoi(timeLimit)
		if err == nil {
			cfg.timeLimit = time.Second * time.Duration(val)
		}

	}

	return cfg
}
