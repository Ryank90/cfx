package exchange

type kraken struct {
}

func krakenExchange(publicKey string, secretKey string, depositAddrs map[string]string) Exchange {
	return &kraken{}
}

func (w *kraken) Name() string {
	return "Kraken"
}

func (w *kraken) String() string {
	return w.Name()
}
