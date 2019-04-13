package exchange

import (
	"github.com/Ryank90/cfx/candle"
	"github.com/Ryank90/cfx/market"
	"github.com/Ryank90/cfx/order"
	"github.com/shopspring/decimal"
)

type simulator struct {
	wrapper  Exchange
	balances map[string]decimal.Decimal
}

func simulatorExchange(w Exchange, fakeBalances map[string]decimal.Decimal) *simulator {
	return &simulator{
		wrapper:  w,
		balances: fakeBalances,
	}
}

// Name does something...
func (w *simulator) Name() string {
	return "Simulator"
}

// String does something...
func (w *simulator) String() string {
	return w.Name()
}

// GetCandles does something...
func (w *simulator) GetCandles(m *market.Market) ([]candle.Stick, error) {

}

// GetOrderBook does something...
func (w *simulator) GetOrderBook(m *market.Market) ([]order.Book, error) {

}

// GetMarketSummary does something...
func (w *simulator) GetMarketSummary(m *market.Market) (*market.Summary, error) {

}

// BuyLimit does something...
func (w *simulator) BuyLimit(m *market.Market, a float64, l float64) (string, error) {

}

// SellLimit does something...
func (w *simulator) SellLimit(m *market.Market, a float64, l float64) (string, error) {

}

// BuyMarket does something...
func (w *simulator) BuyMarket(m *market.Market, a float64) (string, error) {

}

// SellMarket does something...
func (w *simulator) SellMarket(m *market.Market, a float64) (string, error) {

}

// CalculateTradingFees does something...
func (w *simulator) CalculateTradingFees(m *market.Market, a float64, l float64, ot TradeType) float64 {

}

// CalculateWithdrawFees does something...
func (w *simulator) CalculateWithdrawFees(m *market.Market, a float64) float64 {

}

// GetBalance does something...
func (w *simulator) GetBalance(s string) (*decimal.Decimal, error) {

}

// GetDepositAddress does something...
func (w *simulator) GetDepositAddress(c string) (string, bool) {

}

// Withdraw does something...
func (w *simulator) Withdraw(addr string, c string, a float64) error {

}

// Connect connects to the exchange markets.
func (w *simulator) Connect(m []*market.Market) error {
	return w.wrapper.Connect(m)
}
