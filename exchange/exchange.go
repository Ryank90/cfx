package exchange

// Exchange creates an interface for integrating with crypto exchanges.
type Exchange interface {
	// Gets the name of the exchange.
	Name() string
	// Returns a string version of an object.
	String() string
}
