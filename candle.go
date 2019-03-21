package main

import (
	"time"

	"github.com/shopspring/decimal"
)

type candleStick struct {
	// Represents the highest value obtained during candle period.
	High decimal.Decimal
	//
	Open decimal.Decimal
	//
	Close decimal.Decimal
	//
	Low decimal.Decimal
	//
	Volume decimal.Decimal
}

type candleStickChart struct {
	//
	CandlePeriod time.Duration
	//
	CandleSticks []candleStick
	//
	OrderBook []order
}
