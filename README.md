# cryptobot


### Golang Crypto Trading Bot
A golang implementation of a console-based trading bot for cryptocurrency exchanges.


### Usage
Download a release or directly build the code from this repository.
go get github.com/rezhnncoding/cryptobot
If you need to, you can create a strategy and bind it to the bot:

import bot "github.com/rezhnncoding/cryptobot"
func main() {
    bot.AddCustomStrategy(examples.MyStrategy)
    bot.Execute()
}


### Configuration file template
Create a configuration file from this example or run the init command of the compiled executable.
simulation_mode: true # if you want to enable simulation mode.
exchange_configs:
  - exchange: bitfinex
    public_key: bitfinex_public_key
    secret_key: bitfinex_secret_key
    deposit_addresses:
      BTC: bitfinex_deposit_address_btc
      ETH: bitfinex_deposit_address_eth
      ZEC: bitfinex_deposit_address_zec
    fake_balances: # used only if simulation mode is enabled, can be omitted if not enabled.
      BTC: 100
      ETH: 100
      ZEC: 100
      ETC: 100
  - exchange: hitbtc
    public_key: hitbtc_public_key
    secret_key: hitbtc_secret_key
    deposit_addresses:
      BTC : hitbtc_deposit_address_btc
      ETH: hitbtc_deposit_address_eth
      ZEC: hitbtc_deposit_address_zec
    fake_balances:
      BTC: 100
      ETH: 100
      ZEC: 100
      ETC: 100
strategies:
  - strategy: strategy_name
    markets:
      - market: ETH-BTC
        bindings:
        - exchange: bitfinex
          market_name: ETHBTC
        - exchange: hitbtc
          market_name: ETHBTC
      - market: ZEC-BTC
        bindings:
        - exchange: bitfinex
          market_name: ZECBTC
        - exchange: hitbtc
          market_name: ZECBTC
      - market: ETC-BTC
        bindings:
        - exchange: bitfinex
          market_name: ETCBTC
        - exchange: hitbtc
          market_name: ETCBTC
          
     #### donate
     Feel free to donate:
     SEPAH 5892101440517841
     
