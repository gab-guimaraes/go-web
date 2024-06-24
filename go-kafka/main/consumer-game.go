package main

//based on: https://github.com/confluentinc/confluent-kafka-go

import (
	"fmt"
	"go-kafka/model"
	"time"

	"encoding/json"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func main() {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost",
		"group.id":          "myGroup2",
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		panic(err)
	}

	c.SubscribeTopics([]string{"games", "^aRegex.*[Tt]opic"}, nil)

	run := true

	for run {
		msg, err := c.ReadMessage(time.Second)
		if err == nil {
			fmt.Printf("message on %s: %s\n", msg.TopicPartition, string(msg.Value))

			var game model.Game
			err = json.Unmarshal(msg.Value, &game)
			if err != nil {
				fmt.Printf("Erro ao desserializar mensagem", err)
			} else {
				fmt.Printf("Game desserializado com sucesso", game)
			}

		} else if !err.(kafka.Error).IsTimeout() {
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
		}
	}
	c.Close()

}
