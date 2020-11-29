package afra_instrument_v1

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/encoding/gzip"
	"log"
	"time"
)

const (
	address = "192.168.1.108:8102"
	timeout = 60 // in Second period
)

func Connect() (InstrumentClient, context.Context, *grpc.ClientConn, context.CancelFunc) {

	conn, err := grpc.Dial(address, grpc.WithBlock(), grpc.WithInsecure(), grpc.WithDefaultCallOptions(grpc.UseCompressor(gzip.Name)))
	if err != nil {
		log.Fatalf("did not connect to afra_instrument_server : %v", err)
	}

	client := NewInstrumentClient(conn)
	log.Printf(" Message :: Connection is done ")

	ctx, cancel := context.WithTimeout(context.Background(), timeout*time.Second)
	return client, ctx, conn, cancel

}

func GetInstrumentInfoWithCodeList(codelist []string) *SearchInstrumentInfoResponse {
	client, ctx, conn, cancel := Connect()
	defer conn.Close()
	defer cancel()
	query := &SearchInstrumentInfoQuery{CodeList: codelist}
	response, err := client.SearchInstrumentInfo(ctx, query)
	if err != nil {
		log.Fatalf("could not get Instrument with code list %v", err)
	}
	fmt.Println(response.InstrumentList)

	return response
}
func GetInstrumentInfoWithMarketCode(marketCode Market) *SearchInstrumentInfoResponse {
	client, ctx, conn, cancel := Connect()
	defer conn.Close()
	defer cancel()
	response, err := client.SearchInstrumentInfo(ctx, &SearchInstrumentInfoQuery{MarketList: []Market{marketCode}})
	if err != nil {
		log.Fatalf("could not get Instrument with MarketCode %v", err)
	}
	//fmt.Println(response.InstrumentList)
	return response

}

func GetInstrumentSnapshot(refresh_interval int32, code_list []string) Instrument_GetInstrumentSnapshotClient {
	client, ctx, conn, cancel := Connect()
	defer conn.Close()
	defer cancel()
	query := &InstrumentSnapshotQuery{CodeList: code_list, RefreshInterval: refresh_interval}
	stream, err := client.GetInstrumentSnapshot(ctx, query)

	if err != nil {
		log.Fatalf("Couldnt GET Instrumnet SnapShot")
	}
	//for {
	//	response , err := stream.Recv()
	//	if err!= nil{
	//		log.Fatalf("err occured")
	//	}
	//	fmt.Println(response.Code  , response.LastTradedPrice  , response.BuyOrderList )
	//
	//	time.Sleep(time.Duration(refresh_interval)* time.Millisecond)
	//
	//
	//}
	return stream
}

func GetInstrumentHistoricalPrice(time_scale TimeScaleType, start string, end string, code_list []string) *HistoricalPriceResponse {
	client, ctx, conn, cancel := Connect()
	defer conn.Close()
	defer cancel()

	query := &HistoricalPriceQuery{TimeScale: time_scale , Start: start , End: end , CodeList: code_list}
	response , err := client.GetInstrumentHistoricalPrice(ctx , query)
	if err != nil{
		log.Fatalf("Error in GetInstrumentHistoricalPrice")
	}

	return response

}
