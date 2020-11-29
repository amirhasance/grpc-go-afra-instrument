package pubsub

import (
	"fmt"
	"gopkg.in/redis.v2"
	"log"
	"time"
)

type Subscriber struct {
	pubsub   *redis.PubSub
	channel  string
}

type processFunc func(string, string)

func NewSubscriber(channel string) (*Subscriber, error) {
	var err error

	s := Subscriber{
		pubsub:   Service.client.PubSub(),
		channel:  channel,
	}

	// Subscribe to the channel
	err = s.subscribe()
	if err != nil {
		return nil, err
	}

	//// Listen for messages
	go s.listen()

	return &s, nil
}

func (s *Subscriber) subscribe() error {
	var err error

	err = s.pubsub.Subscribe(s.channel)
	if err != nil {
		log.Println("Error subscribing to channel.")
		return err
	}
	return nil
}

func (s *Subscriber) listen() error {

	for {
		msg, err := s.pubsub.ReceiveTimeout(time.Second)
		if err != nil {

			//log.Print("Error in ReceiveTimeout()", err)
		}

		channel := ""
		payload := ""

		switch m := msg.(type) {
		case *redis.Subscription:
			log.Printf("Subscription Message: %v to channel '%v'. %v total subscriptions.", m.Kind, m.Channel, m.Count)
			continue
		case *redis.Message:
			channel = m.Channel
			payload = m.Payload
			fmt.Println(channel , payload  ,1 )
		case *redis.PMessage:
			channel = m.Channel
			payload = m.Payload
			fmt.Println(channel , payload , 1)

		}


	}
}