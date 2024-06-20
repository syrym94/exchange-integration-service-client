package exchange_integration_service_client

import "github.com/syrym94/exchange-integration-service-client/proto"

type ExchangeClient interface {
	GetTrades(exchange string) ([]*proto.Trade, error)
}
