package kafka

import (
	"context"
	"github.com/segmentio/kafka-go"
	"time"
)

type ProducerConfig struct {
	Topic string
	URL   string
}

type TopicConfig struct {
	Topic             string
	URL               string
	Partitions        int
	ReplicationFactor int
}

type Event struct {
	Timestamp time.Time   `json:"timestamp"`
	Message   interface{} `json:"message"`
}

type Producer struct {
	writer *kafka.Writer
}

func NewProducer(cfg *ProducerConfig) *Producer {
	return &Producer{
		writer: &kafka.Writer{
			Addr:     kafka.TCP(cfg.URL),
			Topic:    cfg.Topic,
			Balancer: &kafka.LeastBytes{},
		},
	}
}

func CreateTopic(cfg *TopicConfig) error {
	dialer := &kafka.Dialer{
		Timeout: 10 * time.Second,
	}

	conn, err := dialer.DialContext(context.Background(), "tcp", cfg.URL)
	if err != nil {
		return err
	}
	defer conn.Close()

	topic := cfg.Topic
	err = conn.CreateTopics(kafka.TopicConfig{
		Topic:             topic,
		NumPartitions:     cfg.Partitions,
		ReplicationFactor: cfg.ReplicationFactor,
	})
	if err != nil {
		return err
	}

	return nil
}

func (p *Producer) Write(ctx context.Context, key, value string) error {
	return p.writer.WriteMessages(ctx, kafka.Message{
		Key:   []byte(key),
		Value: []byte(value),
	})
}

func (p *Producer) Close() error {
	return p.writer.Close()
}
