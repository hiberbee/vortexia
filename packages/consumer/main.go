package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/segmentio/kafka-go/sasl/plain"

	godotenv "github.com/joho/godotenv"
	kafka "github.com/segmentio/kafka-go"
)

func loadEnvVariable(key string) string {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	value, exists := os.LookupEnv(key)
	if !exists {
		log.Fatalf("Environment variable %s not set", key)
	}
	return value
}

func getKafkaReader(kafkaURL, topic, groupID string) *kafka.Reader {
	brokers := strings.Split(kafkaURL, ",")
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:  brokers,
		GroupID:  groupID,
		Topic:    topic,
		MinBytes: 10e3, // 10KB
		MaxBytes: 10e6, // 10MB
		Dialer: &kafka.Dialer{
			SASLMechanism: plain.Mechanism{
				Username: loadEnvVariable("KAFKA_USERNAME"),
				Password: loadEnvVariable("KAFKA_PASSWORD"),
			},
		},
	})
}

func main() {
	// get kafka reader using environment variables.
	kafkaURL := loadEnvVariable("KAFKA_URL")
	topic := "timestamps"
	groupID := ""

	reader := getKafkaReader(kafkaURL, topic, groupID)

	defer reader.Close()

	fmt.Println("start consuming ... !!")
	for {
		m, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Printf("message at topic:%v partition:%v offset:%v	%s = %s\n", m.Topic, m.Partition, m.Offset, string(m.Key), string(m.Value))
	}
}
