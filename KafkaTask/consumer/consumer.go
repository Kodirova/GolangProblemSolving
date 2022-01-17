package consumer

import (
	"KafkaTask/api/database"
	"KafkaTask/api/model"
	"encoding/json"
	"fmt"
	"log"

	"github.com/Shopify/sarama"
)

const tableContactCreation = "CREATE TABLE IF NOT EXISTS contacts(id SERIAL, firstname TEXT NOT NULL, lastname TEXT NOT NULL, phone VARCHAR(13), email text, position text)"

func ConnectConsumer(brockersUrl []string) (sarama.Consumer, error) {
	config := sarama.NewConfig()
	config.Producer.Partitioner = sarama.NewManualPartitioner
	config.Consumer.Return.Errors = true

	conn, err := sarama.NewConsumer(brockersUrl, config)

	if err != nil {
		return nil, fmt.Errorf("Failed to ping connection to kafka: %s", err.Error())
	}
	return conn, nil
}

func CreateContact() error {
	var contact *model.Contact
	topic := "create"
	consumer, err1 := ConnectConsumer([]string{"localhost:9092"})
	if err1 != nil {
		panic(err1)
	}
	partitionConsumer, err1 := consumer.ConsumePartition(topic, 0, sarama.OffsetOldest)
	if err1 != nil {
		panic(err1)
	}
	log.Print("Connected to kafka broker")
	for m := range partitionConsumer.Messages() {
		text := string(m.Value)
		bytes := []byte(text)
		json.Unmarshal(bytes, &contact)
		db := database.ConnectDB()
		res, err1 := db.Exec(tableContactCreation)
		if err1 != nil {
			return err1
		}
		fmt.Println(res)
		sqlStatement := `INSERT INTO contacts(firstname, lastname, phone, email, position) VALUES($1, $2, $3, $4, $5) RETURNING id`
		err := db.QueryRow(sqlStatement, &contact.FirstName, &contact.LastName, &contact.Phone, &contact.Email, &contact.Position).Scan(&contact.ID)

		fmt.Println(res)
		if err != nil {
			return err
		}
	}

	return nil
}
