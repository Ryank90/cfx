package candle

import (
	"time"

	"github.com/Ryank90/cfx/order"
	"github.com/shopspring/decimal"
)

// Stick does something.
type Stick struct {
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

// StickChart does something.
type StickChart struct {
	//
	CandlePeriod time.Duration
	//
	CandleSticks []Stick
	//
	OrderBook []order.Order
}
