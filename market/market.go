package market

import "github.com/shopspring/decimal"

// Market does something...
type Market struct {
	//
	Name string `json:"name,required"`
	//
	BaseCurrency string `json:"baseCurrency,omitempty"`
	//
	MarketCurrency string `json:"marketCurrency,omitempty"`
	//
	ExchangeName map[string]string `json:"-"`
}

// Summary does something...
type Summary struct {
	//
	High decimal.Decimal `json:""`
	//
	Low decimal.Decimal `json:""`
	//
	Volume decimal.Decimal `json:""`
	//
	Ask decimal.Decimal `json:""`
	//
	Bid decimal.Decimal `json:""`
	//
	Last decimal.Decimal `json:""`
}

// Ticker does something...
type Ticker struct {
	//
	Ask decimal.Decimal `json:"Ask"`
	//
	Bid decimal.Decimal `json:"Bid"`
	//
	Last decimal.Decimal `json:"Last"`
}

// UpdateFromTicker does something...
func (m *Summary) UpdateFromTicker(t Ticker) {
	m.Ask = t.Ask
	m.Bid = t.Bid
	m.Last = t.Last
}
