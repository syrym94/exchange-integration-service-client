package grpc

import (
	"context"
	"github.com/syrym94/exchange-integration-service-client/proto"
	"google.golang.org/grpc"
)

type ExchangeClient struct {
	client proto.ExchangeServiceClient
}

func NewExchangeClient(conn *grpc.ClientConn) *ExchangeClient {
	c := new(ExchangeClient)

	c.client = proto.NewExchangeServiceClient(conn)

	return c
}

func (ec *ExchangeClient) GetTrades(exchange string) ([]*proto.Trade, error) {
	req := &proto.GetTradesRequest{Exchange: exchange}
	res, err := ec.client.GetTrades(context.Background(), req)
	if err != nil {
		return nil, err
	}
	return res.Trades, nil
}

func (ec *ExchangeClient) GetWalletBalance(exchange, accountType string) (*proto.WalletBalanceResponse, error) {
	req := &proto.GetWalletBalanceRequest{AccountType: accountType, Exchange: exchange}
	res, err := ec.client.GetWalletBalance(context.Background(), req)
	if err != nil {
		return nil, err
	}
	return res.Balance, nil
}

func (ec *ExchangeClient) StreamTicker(exchange, tickerSymbol string) (*proto.TickerResponse, error) {
	req := &proto.TickerRequest{Exchange: exchange, TickerSymbol: tickerSymbol}
	res, err := ec.client.StreamTickerData(context.Background(), req)
	if err != nil {
		return nil, err
	}
	return res.Recv()
}
