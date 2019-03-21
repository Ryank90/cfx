package exchange

import (
	"cfx/config"

	"github.com/shopspring/decimal"
)

// Exchange creates an interface for integrating with crypto exchanges.
type Exchange interface {
	// Gets the name of the exchange.
	Name() string
	// Returns a string version of an object.
	String() string
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
