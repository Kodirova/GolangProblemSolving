package consumer

import (
	"KafkaTask/api/middlewares"
	"KafkaTask/api/model"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/signal"
	_ "strconv"
	"syscall"

	"github.com/Shopify/sarama"
)

const tableContactCreation = "CREATE TABLE IF NOT EXISTS contacts(id SERIAL, firstname TEXT NOT NULL, lastname TEXT NOT NULL, phone VARCHAR(13), email text, position text)"

var DB *sql.DB

func ConnectConsumer(brockersUrl []string) (sarama.Consumer, error) {
	sarama.Logger = log.New(os.Stdout, "[sarama_consumer] ", log.LstdFlags)
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true

	conn, err := sarama.NewConsumer(brockersUrl, config)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func CreateContact() error {
	brokersUrl := []string{"localhost:9092"}
	topic := "create"
	worker, err := ConnectConsumer(brokersUrl)
	if err != nil {
		return err
	}
	consumer, err := worker.ConsumePartition(topic, 0, sarama.OffsetOldest)
	if err != nil {
		return err
	}

	fmt.Println("Consumer started ")
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)
	// Count how many message processed
	msgCount := 0
	var contact model.Contact
	// Get signal for finish
	doneCh := make(chan struct{})
	go func() {
		for {
			select {
			case err := <-consumer.Errors():
				fmt.Println(err)
			case msg := <-consumer.Messages():
				msgCount++
				fmt.Printf("Received message Count %d: | Topic(%s) | Message(%s) \n", msgCount, string(msg.Topic), string(msg.Value))

				middlewares.CreateContact(&contact)
			case <-sigchan:
				fmt.Println("Interrupt is detected")
				doneCh <- struct{}{}
			}
		}
	}()

	<-doneCh
	fmt.Println("Processed", msgCount, "messages")

	if err := worker.Close(); err != nil {
		return err
	}

	return nil
}

func UpdateContact() error {
	brokersUrl := []string{"localhost:9092"}
	topic := "update"
	worker, err := ConnectConsumer(brokersUrl)
	if err != nil {
		panic(err)
	}
	consumer, err := worker.ConsumePartition(topic, 0, sarama.OffsetOldest)
	if err != nil {
		panic(err)
	}
	fmt.Println("Consumer started ")
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)
	// Count how many message processed
	msgCount := 0

	// Get signal for finish
	doneCh := make(chan struct{})
	go func() {
		for {
			select {
			case err := <-consumer.Errors():
				fmt.Println(err)
			case msg := <-consumer.Messages():
				msgCount++
				fmt.Printf("Received message Count %d: | Topic(%s) | Message(%s) | KEY(%s) \n", msgCount, string(msg.Topic), string(msg.Value), msg.Key)
				text := string(msg.Value)
				bytes := []byte(text)
				var contact model.Contact
				log.Println(bytes)
				log.Println(json.Unmarshal(bytes, &contact))
				log.Print("raw : ", text)
				log.Print("user id : ", contact.ID)

				middlewares.UpdateContact(&contact, contact.ID)
			case <-sigchan:
				fmt.Println("Interrupt is detected")
				doneCh <- struct{}{}
			}
		}
	}()

	<-doneCh
	fmt.Println("Processed", msgCount, "messages")

	if err := worker.Close(); err != nil {
		panic(err)
	}

	return nil
}

func DeleteContact() error {
	brokersUrl := []string{"localhost:9092"}
	topic := "delete"
	worker, err := ConnectConsumer(brokersUrl)
	if err != nil {
		panic(err)
	}
	consumer, err := worker.ConsumePartition(topic, 0, sarama.OffsetOldest)
	if err != nil {
		panic(err)
	}
	fmt.Println("Consumer started ")
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)
	// Count how many message processed
	msgCount := 0

	// Get signal for finish
	doneCh := make(chan struct{})
	go func() {
		for {
			select {
			case err := <-consumer.Errors():
				fmt.Println(err)
			case msg := <-consumer.Messages():
				msgCount++
				fmt.Printf("Received message Count %d: | Topic(%s) | Message(%s) | KEY(%s) \n", msgCount, string(msg.Topic), string(msg.Value), msg.Key)
				text := string(msg.Value)
				bytes := []byte(text)
				var contact model.Contact
				log.Println(bytes)
				log.Println(json.Unmarshal(bytes, &contact))
				log.Print("raw : ", text)
				log.Print("user id : ", contact.ID)

				middlewares.DeleteContact(&contact, contact.ID)
			case <-sigchan:
				fmt.Println("Interrupt is detected")
				doneCh <- struct{}{}
			}
		}
	}()

	<-doneCh
	fmt.Println("Processed", msgCount, "messages")

	if err := worker.Close(); err != nil {
		panic(err)
	}

	return nil
}
