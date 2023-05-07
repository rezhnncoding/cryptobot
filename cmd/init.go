package cmd

import (
	"fmt"
	"github.com/saniales/golang-crypto-trading-bot/environment"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initializes the bot to trade",
	Long: `Initializes the trading bot: it will ask several questions to properly create a conf file.
	It must be run prior any other command if config file is not present.`,
	Run: executeInitCommand,
}

func init() {
	RootCmd.AddCommand(initCmd)
	initCmd.Flags().StringVar(&initFlags.ConfigFile, "import", "", "imports configuration from a file.")
}

func executeInitCommand(cmd *cobra.Command, args []string) {
	initConfig()
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if initFlags.ConfigFile != "" {
		//try first to unmarshal the file to check if it is correct format.
		content, err := ioutil.ReadFile(initFlags.ConfigFile)
		if err != nil {
			fmt.Print("Error while opening the config file provided")
			if GlobalFlags.Verbose > 0 {
				fmt.Printf(": %s", err.Error())
			}
			fmt.Println()
			return
		}
		var checker environment.BotConfig
		err = yaml.Unmarshal(content, &checker)
		if err != nil {
			fmt.Print("Cannot load provided configuration file")
			if GlobalFlags.Verbose > 0 {
				fmt.Printf(": %s", err.Error())
			}
			fmt.Println()
			return
		}
		err = ioutil.WriteFile("./.bot_config.yml", content, 0666)
		if err != nil {
			fmt.Print("Cannot write new configuration file")
			if GlobalFlags.Verbose > 0 {
				fmt.Printf(": %s", err.Error())
			}
			fmt.Println()
			return
		}
	} else {
		generateInitFile()
	}
}

func generateInitFile() {
	configs := environment.BotConfig{}
	for {
		var exchange environment.ExchangeConfig
		var YesNo string

		fmt.Println("Which exchange are you going to add?")
		fmt.Scanln(&exchange.ExchangeName)

		alreadyAdded := false
		for _, ex := range configs.ExchangeConfigs {
			if ex.ExchangeName == exchange.ExchangeName {
				alreadyAdded = true
				break
			}
		}

		if alreadyAdded {
			fmt.Println("Exchange already added, retry.")
			continue
		}

		fmt.Println("Please provide Public Key for that exchange.")
		fmt.Scanln(&exchange.PublicKey)
		fmt.Println("Please provide Secret Key for that exchange.")
		fmt.Scanln(&exchange.SecretKey)

		configs.ExchangeConfigs = append(configs.ExchangeConfigs, exchange)

		fmt.Println("Exchange Added")
		for YesNo != "Y" && YesNo != "n" {
			fmt.Println("Do you want to add another exchange? (Y/n)")
			fmt.Scanln(&YesNo)
		}
		if YesNo == "n" {
			break
		}
	}

	for {
		var YesNo string
		for YesNo != "Y" && YesNo != "n" {
			fmt.Println("Do you want to add a strategy binded to a market? (Y/n)")
			fmt.Scanln(&YesNo)
		}
		if YesNo == "n" {
			break
		}

		var tempStrategyAppliance environment.StrategyConfig

		fmt.Println("Please Enter The Name of the strategy you want to use\n" +
			"in this market (must be in the system)")
		fmt.Scanln(&tempStrategyAppliance.Strategy)

		for {
			var tmpMarketConf environment.MarketConfig
			fmt.Println("Please Enter Market Name using short notation " +
				"(e.g. BTC-ETH for a Bitcoin-Ethereum market).")
			fmt.Scanln(&tmpMarketConf.Name)
			for _, ex := range configs.ExchangeConfigs {
				var exMarketName string
				fmt.Printf("Please Enter %s exchange market ticker, or leave empty to skip this exchange\n", ex.ExchangeName)
				fmt.Scanln(&exMarketName)

				if exMarketName != "" {
					tmpMarketConf.Exchanges = append(tmpMarketConf.Exchanges, environment.ExchangeBindingsConfig{
						Name:       ex.ExchangeName,
						MarketName: exMarketName,
					})
					fmt.Printf("Exchange %s CONFIGURED with Market Name %s\n", ex.ExchangeName, exMarketName)
				} else {
					fmt.Printf("Exchange %s SKIPPED\n", ex.ExchangeName)
				}
			}

			tempStrategyAppliance.Markets = append(tempStrategyAppliance.Markets, tmpMarketConf)

			var YesNo string
			for YesNo != "Y" && YesNo != "n" {
				fmt.Println("Do you want to add another market binded to this strategy? (Y/n)")
				fmt.Scanln(&YesNo)
			}
			if YesNo == "n" {
				break
			}
		}

		configs.Strategies = append(configs.Strategies, tempStrategyAppliance)
	}

	//preview the contents of the file to be written, then creates a new file in .
	contentToBeWritten, err := yaml.Marshal(configs)
	if err != nil {
		fmt.Print("Error while creating the content for the new config file")
		if GlobalFlags.Verbose > 0 {
			fmt.Printf(": %s", err.Error())
		}
		fmt.Println()
		return
	}
	fmt.Println("The following content:")
	fmt.Println(string(contentToBeWritten))
	fmt.Println("is going to be written on ./.gobot, is it ok? (Y/n)")

	var YesNo string
	for YesNo != "Y" && YesNo != "n" {
		fmt.Scanln(&YesNo)
	}
	if YesNo == "Y" {
		err := ioutil.WriteFile("./.gobot", contentToBeWritten, 0666)
		if err != nil {
			fmt.Print("Error while writing content to new config file")
			if GlobalFlags.Verbose > 0 {
				fmt.Printf(": %s", err.Error())
			}
			fmt.Println()
		} else {
			fmt.Println("Config file created on ./.gobot\nNow you can use gobot with this configuration.\nHappy Trading, folk :)")
		}
		return
	}
	fmt.Println("You chose not to write the content to configuration file.\n" +
		"You can relaunch this command again to create another configuration.\n" +
		"This bot won't work until it has a valid configuration file.")
}
