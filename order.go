package main

import (
	"time"

	"github.com/shopspring/decimal"
)

// orderType is an enum.
type orderType int16

const (
	// Ask represents an Ask order.
	Ask orderType = iota
	// Bid represents a Bid order.
	Bid orderType = iota
)

// order represents a single order within the orderBook for a market.
type order struct {
	// Value of the trade.
	Value decimal.Decimal
	// Quantity of coins for this order.
	Qty decimal.Decimal
	// Order number as seen in the exchange.
	OrderNumber string
	// Timestamp of when the order was made.
	Timestamp time.Time
}

// orderBook represents a standard order book implementation.
type orderBook struct {
	Asks []order `json:"asks,required"`
	Bids []order `json:"bids,required"`
}

// Returns the order total in the base currency.
func (o order) total() decimal.Decimal {
	return o.Qty.Mul(o.Value)
}
