package kafka

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/rohanchauhan02/recommendation-engine/pkg/config"
)

type IKafka interface {
	Init() (*kafka.Producer, error)
}

type KafkaClient struct {
	conf config.ImmutableConfig
}

