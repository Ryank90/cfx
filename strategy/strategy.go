package strategy

import (
    "cfx/exchange"
    "cfx/market"
)

// Strategy creates an interface for setting up strategies.
type Strategy interface {
    // Returns the name of the strategy.
    Name() string
    // Applies a strategy when called using the specified wrapper.
    Apply([]exchange.Exchange, []*market.Market)
}
