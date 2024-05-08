package main

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

type data struct {
	X uint64 `json:"x"`
	Y uint64 `json:"y"`
}

func main() {
	// to produce messages
	topic := "notifications"
	partition := 0

	conn, err := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", topic, partition)
	if err != nil {
		log.Fatal("failed to dial leader:", err)
	}

	conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
	for i := uint64(0); i < 1000; i++ {
		d := data{
			X: i * 2,
			Y: i * i,
		}
		input, _ := json.Marshal(d)
		_, err = conn.WriteMessages(
			kafka.Message{Value: input},
		)
		if err != nil {
			log.Fatal("failed to write messages:", err)
		}
	}

	if err := conn.Close(); err != nil {
		log.Fatal("failed to close writer:", err)
	}
	// r := kafka.NewReader(kafka.ReaderConfig{
	// 	Brokers:  []string{"localhost:9092"},
	// 	GroupID:  "vector_consumer_group_xxx",
	// 	Topic:    "your_topic_here",
	// 	MaxBytes: 10e6, // 10MB
	// })

	// for {
	// 	m, err := r.ReadMessage(context.Background())
	// 	if err != nil {
	// 		break
	// 	}
	// 	fmt.Printf("message at topic/partition/offset %v/%v/%v: %s = %s\n", m.Topic, m.Partition, m.Offset, string(m.Key), string(m.Value))
	// }

	// if err := r.Close(); err != nil {
	// 	log.Fatal("failed to close reader:", err)
	// }
}
