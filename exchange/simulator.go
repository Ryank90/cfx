package exchange

import "github.com/shopspring/decimal"

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

func (w *simulator) Name() string {
	return "Simulator"
}

func (w *simulator) String() string {
	return w.Name()
}
