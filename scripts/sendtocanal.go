package main

import (
	"github.com/nats-io/stan.go"
	"log"
	"os"
)

type natsConfig struct {
	NatsURL   string
	ClusterID string
	ClientID  string
}

func main() {

	file, err := os.ReadFile("scripts/data.json") //невалидная модель datainvalid.json, валидная data.json
	if err != nil {
		log.Fatal(err)
	}

	// можно реализовать проверку на соответствие Order, для того, чтобы быть уверенным
	// что в канал попадут нужные данные, но в данных условиях , будем кидать все что душе угодно
	// для проверки consumer-а

	config := natsConfig{
		NatsURL:   "nats://localhost:4222",
		ClusterID: "test-cluster",
		ClientID:  "sender",
	}

	nats, err := stan.Connect(config.ClusterID, config.ClientID, stan.NatsURL(config.NatsURL))
	if err != nil {
		log.Fatal(err)
	}
	defer nats.Close()

	err = nats.Publish("foo", file)
	if err != nil {
		log.Fatal(err)
	}

	//sub, err := nats.Subscribe("foo", func(m *stan.Msg) {
	//	fmt.Println("from nats")
	//	fmt.Println(string(m.Data))
	//}, stan.DurableName("my-durable"))
	//if err != nil {
	//	log.Fatal(err)
	//}

	//defer sub.Unsubscribe()
}
