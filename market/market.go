package market

import "github.com/shopspring/decimal"

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

type MarketSummary struct {
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
