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
	// Gets the candle data from the exchange.
	GetCandles(m *market.Market) ([]candle.Stick, error)
	// Gets the order book from the exchange.
	GetOrderBook(m *market.Market) ([]order.Book, error)
	// Gets the market summary from the exchange.
	GetMarketSummary(m *market.Market) ([]market.Summary, error)
	// Performs a limit buy action.
	BuyLimit(m *market.Market, amount float64, limit float64) (string, error)
	// Performs a limit sell action.
	SellLimit(m *market.Market, amount float64, limit float64) (string, error)
	// Performs a buy action on a specific market.
	BuyMarket(m *market.Market, amount float64) (string, error)
	// Performs a sell action on a specific market.
	SellMarket(m *market.Market, amount float64) (string, error)
	// Calculate the trading fees for a specific action.
	CalculateTradingFees(m *market.Market, amount float64, limit float64, orderType TradeType) float64
	// Calculate the withdrawal fee for a specific action.
	CalculateWithdrawFees(m *market.Market, amount float64) float64
	// Gets a balance for the specified user.
	GetBalance(s string) (*decimal.Decimal, error)
	// Gets the deposit address for the specified coin on the exchange if it is present.
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
