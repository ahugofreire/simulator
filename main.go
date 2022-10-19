package main

import (
	"fmt"
	"log"

	"github.com/ahugofreire/simulator/infra/kafka"
	"github.com/joho/godotenv"

	kafka2 "github.com/ahugofreire/simulator/application/kafka"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}
}

func main() {

	// producer := kafka.NewKafkaProducer()
	// kafka.Publish("test", "readtest", producer)

	// for {
	// 	_ = 1
	// }

	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatal("error loading .env file")
	// }
	// route := route.Route{
	// 	ID:       "1",
	// 	ClientID: "1",
	// }
	// route.LoadPositions()
	// strJson, _ := route.ExportJsonPositions()
	// fmt.Println(strJson[1])

	//msg_sample {"clientId":"1","routeId":"1"}
	//msg_sample {"clientId":"2","routeId":"2"}
	msgChan := make(chan *ckafka.Message)
	consumer := kafka.NewKafkaConsumer(msgChan)
	go consumer.Consume()

	for msg := range msgChan {
		fmt.Println(string(msg.Value))
		go kafka2.Produce(msg)
	}
}
