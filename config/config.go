package config

import "github.com/shopspring/decimal"

// Config stores the main configuration for the bot.
type Config struct {
	// Are we running the bot in simulation mode?
	SimulationMode bool `yaml:"simulation_mode"`
	// List of exchanges that are registered with the bot.
	Exchanges []ExchangeConfig `yaml:"exchanges"`
	// List of strategies that are being used by the bot.
	Strategies []StrategyConfig `yaml:"strategies"`
}

// ExchangeConfig stores exchange specific configration.
type ExchangeConfig struct {
	// Name of the exchange.
	ExchangeName string `yaml:"exchange_name"`
	// URI of the exchange.
	ExchangeURI string `yaml:"exchange_uri"`
	// Public key for interfacing with the exchange API.
	PublicKey string `yaml:"public_key"`
	// Secret key for interfacing with the exchange API.
	SecretKey string `yaml:"secret_key"`
	// Direct connection between desposit addresses and coins on the exchange.
	DepositAddress map[string]string `yaml:"deposit_addresses"`
	// Test balances for coins.
	TestBalances map[string]decimal.Decimal `yaml:"test_balances"`
}

// StrategyConfig stores strategy specific configuration.
type StrategyConfig struct {
	// Strategy that is going to be used with the bot.
	Strategy string `yaml:"strategy"`
	// Markets in-which the strategy will impact.
	Markets []MarketConfig `yaml:"markets"`
}

// MarketConfig stores market specific configuration.
type MarketConfig struct {
	// Market name.
	Name string `yaml:"market"`
}
