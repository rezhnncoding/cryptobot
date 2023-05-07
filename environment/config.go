package environment

import "github.com/shopspring/decimal"

// ExchangeConfig Represents a configuration for an API Connection to an exchange.
//
//	Can be used to generate an ExchangeWrapper.
type ExchangeConfig struct {
	ExchangeName     string                     `yaml:"exchange"`          // Represents the exchange name.
	PublicKey        string                     `yaml:"public_key"`        // Represents the public key used to connect to Exchange API.
	SecretKey        string                     `yaml:"secret_key"`        // Represents the secret key used to connect to Exchange API.
	DepositAddresses map[string]string          `yaml:"deposit_addresses"` // Represents the bindings between coins and deposit address on the exchange.
	FakeBalances     map[string]decimal.Decimal `yaml:"fake_balances"`     // Used only in simulation mode, fake starting balance [coin:balance].
}

// StrategyConfig contains where a strategy will be applied in the specified exchange.
type StrategyConfig struct {
	Strategy string         `yaml:"strategy"` // Represents the applied strategy name: must be unique in the system.
	Markets  []MarketConfig `yaml:"markets"`  // Represents the exchanges where the strategy is applied.
}

// MarketConfig contains all market configuration data.
type MarketConfig struct {
	Name      string                   `yaml:"market"`   // Represents the market where the strategy is applied.
	Exchanges []ExchangeBindingsConfig `yaml:"bindings"` // Represents the list of markets where the strategy is applied, along with extra-data regarding binded exchanges.
}

// ExchangeBindingsConfig represents the binding of market names between bot notation and exchange ticker.
type ExchangeBindingsConfig struct {
	Name       string `yaml:"exchange"`    // Represents the name of the exchange.
	MarketName string `yaml:"market_name"` // Represents the name of the market as seen from the exchange.
}

// BotConfig contains all config data of the bot, which can be also loaded from config file.
type BotConfig struct {
	SimulationModeOn bool             `yaml:"simulation_mode"`  // if true, do not create real orders and do not get real balance
	ExchangeConfigs  []ExchangeConfig `yaml:"exchange_configs"` // Represents the current exchange configuration.
	Strategies       []StrategyConfig `yaml:"strategies"`       // Represents the current strategies adopted by the bot.
}
