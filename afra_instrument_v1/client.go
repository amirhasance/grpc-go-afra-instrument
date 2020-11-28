package afra_instrument_v1

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/encoding/gzip"
	"log"
	"time"
)

const (
	address            = "192.168.1.108:8102"
	timeout_in_seconds = 60
)

func Connect() ( InstrumentClient, context.Context  , *grpc.ClientConn   , context.CancelFunc) {

	conn, err := grpc.Dial(address, grpc.WithBlock(), grpc.WithInsecure(), grpc.WithDefaultCallOptions(grpc.UseCompressor(gzip.Name)))
	if err != nil {
		log.Fatalf("did not connect to afra_instrument_server : %v", err)
	}

	client := NewInstrumentClient(conn)

	log.Printf(" Message :: Connection is done ")

	ctx, cancel := context.WithTimeout(context.Background(), timeout_in_seconds*time.Second)

	return client, ctx , conn , cancel

}

func GetWithCodeList(codelist []string) ( *SearchInstrumentInfoResponse) {
	client , ctx , conn , cancel := Connect()
	defer conn.Close()
	defer cancel()
	query := &SearchInstrumentInfoQuery{CodeList: codelist}
	response, err := client.SearchInstrumentInfo(ctx, query)
	if err != nil {
		log.Fatalf("could not get Instrument with code list %v", err)
	}


	return response

}

func GetWithMarketCode(marketCode Market) ( *SearchInstrumentInfoResponse ){
	client, ctx  , conn  , cancel:= Connect()
	defer conn.Close()
	defer cancel()

	response, err := client.SearchInstrumentInfo(ctx, &SearchInstrumentInfoQuery{MarketList: []Market{marketCode}})
	if err != nil {
		log.Fatalf("could not get Instrument with MarketCode %v", err)
	}
	return response

}
