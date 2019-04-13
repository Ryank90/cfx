package exchange

import (
	"github.com/Ryank90/cfx/candle"
	"github.com/Ryank90/cfx/config"
	"github.com/Ryank90/cfx/market"
	"github.com/Ryank90/cfx/order"
	"github.com/shopspring/decimal"
)

// TradeType represents an order type.
type TradeType string

// Exchange creates an interface for integrating with crypto exchanges.
type Exchange interface {
	// Gets the name of the exchange.
	Name() string
	// Returns a string version of an object.
	String() string
	//
	GetCandles(m *market.Market) ([]candle.Stick, error)
	//
	GetOrderBook(m *market.Market) ([]order.Book, error)
	//
	GetMarketSummary(m *market.Market) ([]market.Summary, error)
	//
	BuyLimit(m *market.Market, amount float64, limit float64) (string, error)
	//
	SellLimit(m *market.Market, amount float64, limit float64) (string, error)
	//
	BuyMarket(m *market.Market, amount float64) (string, error)
	//
	SellMarket(m *market.Market, amount float64) (string, error)
	//
	CalculateTradingFees()
	//
	CalculateWithdrawFees()
	//
	GetBalance(s string) (*decimal.Decimal, error)
	//
	GetDepositAddress(ct string) (string, bool)
	// Performs a withdrawal from an exchange to the desired deposit address.
	Withdraw(depAddr string, ct string, amount float64) error
	// Connects to the feed of the exchange.
	Connect(m []*market.Market) error
}

// InitExchange initiates a new Exchange binded to the specified exchange provided.
func InitExchange(cfg config.ExchangeConfig, simulationMode bool, fakeBalances map[string]decimal.Decimal, depositAddrs map[string]string) Exchange {
	if depositAddrs == nil && !simulationMode {
		return nil
	}

	var ex Exchange
	switch cfg.ExchangeName {
	case "kraken":
		ex = krakenExchange(cfg.PublicKey, cfg.SecretKey, depositAddrs)
	default:
		return nil
	}

	if simulationMode {
		if fakeBalances == nil {
			return nil
		}

		ex = simulatorExchange(ex, fakeBalances)
	}

	return ex
}

// MarketName gets the name of the market as seen by the exchange.
func MarketName(m *market.Market, w Exchange) string {
	return m.ExchangeName[w.Name()]
}
