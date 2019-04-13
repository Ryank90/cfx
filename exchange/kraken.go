package exchange

import (
	"github.com/Ryank90/cfx/candle"
	"github.com/Ryank90/cfx/market"
	"github.com/Ryank90/cfx/order"
	"github.com/shopspring/decimal"
)

type kraken struct {
}

func krakenExchange(publicKey string, secretKey string, depositAddrs map[string]string) Exchange {
	return &kraken{}
}

// Name does something...
func (w *kraken) Name() string {
	return "Kraken"
}

// String does something...
func (w *kraken) String() string {
	return w.Name()
}

// GetCandles does something...
func (w *kraken) GetCandles(m *market.Market) ([]candle.Stick, error) {

}

// GetOrderBook does something...
func (w *kraken) GetOrderBook(m *market.Market) ([]order.Book, error) {

}

// GetMarketSummary does something...
func (w *kraken) GetMarketSummary(m *market.Market) (*market.Summary, error) {

}

// BuyLimit does something...
func (w *kraken) BuyLimit(m *market.Market, a float64, l float64) (string, error) {

}

// SellLimit does something...
func (w *kraken) SellLimit(m *market.Market, a float64, l float64) (string, error) {

}

// BuyMarket does something...
func (w *kraken) BuyMarket(m *market.Market, a float64) (string, error) {

}

// SellMarket does something...
func (w *kraken) SellMarket(m *market.Market, a float64) (string, error) {

}

// CalculateTradingFees does something...
func (w *kraken) CalculateTradingFees(m *market.Market, a float64, l float64, ot TradeType) float64 {

}

// CalculateWithdrawFees does something...
func (w *kraken) CalculateWithdrawFees(m *market.Market, a float64) float64 {

}

// GetBalance does something...
func (w *kraken) GetBalance(s string) (*decimal.Decimal, error) {

}

// GetDepositAddress does something...
func (w *kraken) GetDepositAddress(c string) (string, bool) {

}

// Withdraw does something...
func (w *kraken) Withdraw(addr string, c string, a float64) error {

}

// Connect connects to the exchange markets.
func (w *kraken) Connect(m []*market.Market) error {
	return w.wrapper.Connect(m)
}
