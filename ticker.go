package main

import "github.com/shopspring/decimal"

type ticker struct {
	//
	Ask decimal.Decimal `json:"Ask"`
	//
	Bid decimal.Decimal `json:"Bid"`
	//
	Last decimal.Decimal `json:"Last"`
}

func (ms *marketSummary) updateFromTicker(t ticker) {
	ms.Ask = t.Ask
	ms.Bid = t.Bid
	ms.Last = t.Last
}
