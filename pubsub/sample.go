package pubsub

import (
	"gopkg.in/redis.v2"
	"log"
	"time"
)

func test_pubsub() {

	type food struct {
		Name     string
		Calories float64
	}

	type car struct {
		Make  string
		Model string
	}

	var pub *redis.IntCmd
	var err error

	// Create a subscriber
	_, err = NewSubscriber("food")
	if err != nil {
		log.Println("NewSubscriber() error", err)
	}

	// Create another subscriber
	_, err = NewSubscriber("cars")
	if err != nil {
		log.Println("NewSubscriber() error", err)
	}
	_, err = NewSubscriber("cars")
	if err != nil {
		log.Println("NewSubscriber() error", err)
	}
	log.Print("Subscriptions done. Publishing...")

	// -- Publish some stuf --
	pub = Service.Publish("food", food{"Pizza", 50.1})
	if err = pub.Err(); err != nil {
		log.Print("PublishString() error", err)
	}

	pub = Service.Publish("food", food{"Big Mac", 200})
	if err = pub.Err(); err != nil {
		log.Print("PublishString() error", err)
	}

	pub = Service.Publish("cars", car{"Subaru", "Impreza"})
	if err = pub.Err(); err != nil {
		log.Print("PublishString() error", err)
	}

	pub = Service.Publish("cars", car{"Tesla", "Model S"})
	if err = pub.Err(); err != nil {
		log.Print("PublishString() error", err)
	}
	_ = pub

	log.Print("Publishing done. Sleeping...")

	time.Sleep(time.Second)
}

