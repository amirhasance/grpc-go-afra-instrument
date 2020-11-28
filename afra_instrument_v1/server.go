package afra_instrument_v1

import (
	"fmt"
	"google.golang.org/grpc"
	"log"
)
const (
	address = "192.168.1.108:8102"

	)


func Connect() {
	conn  , err := grpc.Dial( address , grpc.WithInsecure() , grpc.WithBlock())

	if err != nil {
		log.Fatalf("did not connect : %v" , err)
	}
	defer conn.Close()
	c := NewInstrumentClient(conn)

	fmt.Println(c)

	fmt.Println("connection is done ")
	//ctx , cancel := context.WithTimeout(context.Background() , time.Second)

	//defer cancel()

	//response , err := c.GetInstrumentHistoricalPrice( ctx , $HistoricalPriceQuery{})
	//response , err := c.GetInstrumentSnapshot()
	//
	//response , err := c.SearchInstrumentInfo(ctx , &SearchInstrumentInfoQuery{})
	//
	//if err != nil {
	//	log.Fatalf("could not get Instrument Historical %v" , err)
	//}
	//log.Printf("Message : %s"  , response.InstrumentList)



}