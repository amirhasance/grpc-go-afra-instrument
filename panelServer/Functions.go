package panelServer

import "context"

type Server struct {

}
func (s *Server) SearchInstrumentInfo(ctx context.Context, query *SearchInstrumentInfoQuery) (*SearchInstrumentInfoResponse, error){}
func(s *Server)GetInstrumentSnapshot(query *InstrumentSnapshotQuery, snap Instrument_GetInstrumentSnapshotServer) error{}
func(s *Server)GetInstrumentHistoricalPrice(ctx context.Context, qeury *HistoricalPriceQuery) (*HistoricalPriceResponse, error){}