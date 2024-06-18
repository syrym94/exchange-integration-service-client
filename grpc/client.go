package grpc

import (
	"context"
	"exchange-integration-service-client/entity"
	ExchangeProtobuf "exchange-integration-service-client/protobuf"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

type Client struct {
	address     string
	conn        *grpc.ClientConn
	dialOpts    []grpc.DialOption
	middlewares []grpc.UnaryClientInterceptor
}

func NewClient() (*Client, error) {
	cli := &Client{
		middlewares: make([]grpc.UnaryClientInterceptor, 0),
		dialOpts:    make([]grpc.DialOption, 0),
	}

	if err := cli.connect(); err != nil {
		return nil, err
	}

	return cli, nil
}

func (c *Client) connect() error {
	if len(c.middlewares) == 1 {
		c.dialOpts = append(c.dialOpts, grpc.WithUnaryInterceptor(c.middlewares[0]))
	} else if len(c.middlewares) > 1 {
		c.dialOpts = append(c.dialOpts, grpc.WithChainUnaryInterceptor(c.middlewares...))
	}

	conn, err := grpc.Dial(c.address, c.dialOpts...)
	if err != nil {
		return errors.Wrap(err, "[ERROR] Error establishing protobuf connection to jam seo")
	}

	c.conn = conn

	return nil
}

func (c *Client) Client() ExchangeProtobuf.ExchangeServiceClient {
	return ExchangeProtobuf.NewExchangeServiceClient(c.conn)
}

func (c *Client) GetTrades(ctx context.Context, resp entity.Trade) (*ExchangeProtobuf.GetTradesResponse, error) {
	return c.Client().GetTrades(ctx, resp.TradeReqToProto())
}

func (c *Client) Close() error {
	return c.conn.Close()
}
