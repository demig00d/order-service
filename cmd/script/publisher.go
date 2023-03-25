package main

import (
	"encoding/json"
	"fmt"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/demig00d/order-service/config"
	"github.com/demig00d/order-service/internal/entity"
	"github.com/demig00d/order-service/pkg/jetstream"
	"log"
	"math/rand"
	"time"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	log.Println("Starting...")

	// Connect to NATS
	js, err := jetstream.New(cfg.JetStream)
	if err != nil {
		log.Fatal(fmt.Errorf("script - Run - jetstream.New: %w", err))
	}
	defer js.Close()

	err = js.CreateStream()

	if err != nil {
		log.Fatal(fmt.Errorf("script - Run - js.CreateStream(): %w", err))
	}

	publishOrders(js)
}

func publishOrders(js *jetstream.JetStream) {

	for {
		order, err := getOrder()

		if err != nil {
			log.Println(err)
			continue
		}

		orderJson, err := json.Marshal(order)
		if err != nil {
			log.Println(err)
			continue
		}

		// publish to ORDERS.received subject
		_, err = js.Publish(js.SubjectNameOrderCreated, orderJson)
		if err != nil {
			log.Println(err)
		} else {
			log.Printf("Publisher  =>  OrderUid:%s\n", order.OrderUid)
		}
		// create random message intervals to slow down
		r := rand.Intn(1500)
		time.Sleep(time.Duration(r) * time.Millisecond)
	}
}

func getOrder() (entity.Order, error) {
	var order entity.Order

	err := gofakeit.Struct(&order)

	return order, err
}
