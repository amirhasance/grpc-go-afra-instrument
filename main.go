package main

import (
	"gopkg.in/redis.v2"
	"grpcPanel/pubsub"
	"log"
	"time"
)

func main() {
	//afra_instrument_v1.GetInstrumentInfoWithCodeList([]string{"8516759016718718"})
	//af.GetInstrumentInfoWithCodeList([]string{"8516759016718718", "13082825954602280", "48711126865295010"})
	//af.GetWithMarketCode(2)
	//afra_instrument_v1.GetInstrumentSnapshot(1000 ,   []string{"13082825954602280" , "48711126865295010" , "8516759016718718"})
	//afra_instrument_v1.GetInstrumentHistoricalPrice(afra_instrument_v1.TimeScaleType_Day_1 , "2020-10-10 11:00:00" , "2020-10-13 11:00:00"  , []string{"8516759016718718"} )
	//test_pubsub()

}


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
	_, err = pubsub.NewSubscriber("food")
	if err != nil {
		log.Println("NewSubscriber() error", err)
	}

	// Create another subscriber
	_, err = pubsub.NewSubscriber("cars")
	if err != nil {
		log.Println("NewSubscriber() error", err)
	}
	_, err = pubsub.NewSubscriber("cars")
	if err != nil {
		log.Println("NewSubscriber() error", err)
	}
	log.Print("Subscriptions done. Publishing...")

	// -- Publish some stuf --
	pub = pubsub.Service.Publish("food", food{"Pizza", 50.1})
	if err = pub.Err(); err != nil {
		log.Print("PublishString() error", err)
	}

	pub = pubsub.Service.Publish("food", food{"Big Mac", 200})
	if err = pub.Err(); err != nil {
		log.Print("PublishString() error", err)
	}

	pub = pubsub.Service.Publish("cars", car{"Subaru", "Impreza"})
	if err = pub.Err(); err != nil {
		log.Print("PublishString() error", err)
	}

	pub = pubsub.Service.Publish("cars", car{"Tesla", "Model S"})
	if err = pub.Err(); err != nil {
		log.Print("PublishString() error", err)
	}
	_ = pub

	log.Print("Publishing done. Sleeping...")

	time.Sleep(time.Second)
}




