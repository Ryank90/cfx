package order

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

// Order represents a single order within the orderBook for a market.
type Order struct {
	// Value of the trade.
	Value decimal.Decimal
	// Quantity of coins for this order.
	Qty decimal.Decimal
	// Order number as seen in the exchange.
	OrderNumber string
	// Timestamp of when the order was made.
	Timestamp time.Time
}

// Book represents a standard order book implementation.
type Book struct {
	Asks []Order `json:"asks,required"`
	Bids []Order `json:"bids,required"`
}

// Returns the order total in the base currency.
func (o Order) total() decimal.Decimal {
	return o.Qty.Mul(o.Value)
}
