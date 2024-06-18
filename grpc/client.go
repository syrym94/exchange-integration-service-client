package clients

import (
	"context"
	proto "github.com/syrym94/exchange-integration-service-client/protobuf"
	"google.golang.org/grpc"
)

type ExchangeClient struct {
	client proto.ExchangeServiceClient
}

func NewExchangeClient(address string) (*ExchangeClient, error) {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	client := proto.NewExchangeServiceClient(conn)
	return &ExchangeClient{client: client}, nil
}

func (ec *ExchangeClient) GetTrades(exchange string) ([]*proto.Trade, error) {
	req := &proto.GetTradesRequest{Exchange: exchange}
	res, err := ec.client.GetTrades(context.Background(), req)
	if err != nil {
		return nil, err
	}
	return res.Trades, nil
}
