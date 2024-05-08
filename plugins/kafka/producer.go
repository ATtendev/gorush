package kafkax

import (
	"context"
	"time"

	"github.com/appleboy/gorush/config"
	"github.com/segmentio/kafka-go"
)

var Producer *producer

type producer struct {
	ctx    context.Context
	config config.Kafka
	conn   *kafka.Conn
}

func NewProducer(config *config.Kafka) {
	ctx := context.Background()
	dialer := &kafka.Dialer{
		Timeout:   10 * time.Second,
		DualStack: true,
	}
	conn, err := dialer.DialLeader(ctx, config.Network, config.Address, config.Topic, 0)
	if err != nil {
		panic(err.Error())
	}
	Producer = &producer{
		ctx:    ctx,
		config: *config,
		conn:   conn,
	}
}

func (p *producer) SendMessage(data []byte) error {
	_, err := p.conn.WriteMessages(kafka.Message{
		Value: data,
	})
	return err
}

func (p *producer) Close() error {
	return p.conn.Close()
}
