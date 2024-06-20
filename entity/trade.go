package entity

import (
	"github.com/syrym94/exchange-integration-service-client/proto"
)

type Trade struct {
	Symbol   string
	Price    float64
	Quantity float64
	ID       string
	Time     int64
}

func (Trade) TradeReqToProto() *proto.GetTradesRequest {
	return &proto.GetTradesRequest{
		Exchange: "",
	}
}

func (Trade) TradeToProto(t *Trade) *proto.Trade {
	return &proto.Trade{
		Id:       t.ID,
		Price:    t.Price,
		Quantity: t.Quantity,
		Symbol:   t.Symbol,
		Time:     t.Time,
	}
}

func (Trade) TradeFromProto(t *proto.Trade) *Trade {
	return &Trade{
		Symbol:   t.Symbol,
		Price:    t.Price,
		Quantity: t.Quantity,
		ID:       t.Id,
		Time:     t.Time,
	}
}
