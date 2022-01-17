package producer

import (
	"fmt"
	"log"
	"os"

	"github.com/Shopify/sarama"
)

func ConnectProducer(brokersUrl []string) (sarama.SyncProducer, error) {
	sarama.Logger = log.New(os.Stdout, "[sarama] ", log.LstdFlags)
	config := sarama.NewConfig()
	config.Version = sarama.V3_0_0_0
	config.Producer.Return.Successes = true
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5
	conn, err := sarama.NewSyncProducer(brokersUrl, config)

	if err != nil {
		panic(err)
	}
	return conn, nil
}

func PushHandlerToQueue(topic string, message []byte) error {
	brokersUrl := []string{"localhost:9092"}
	producer, err := ConnectProducer(brokersUrl)
	if err != nil {
		return err
	}
	data := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(message),
	}
	fmt.Println(topic)
	partition, offset, err := producer.SendMessage(data)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Message is stored in topic (%s)/partition(%d)/offset(%d)\n", topic, partition, offset)
	return nil
}
