package exchange

type simulator struct {
}

func simulatorExchange() Exchange {
	return &simulator{}
}

func (w *simulator) Name() string {
	return "Simulator"
}

func (w *simulator) String() string {
	return w.Name()
}
