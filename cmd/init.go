package cfx

import (
	"fmt"
	"io/ioutil"

	"github.com/Ryank90/cfx/config"
	yaml "gopkg.in/yaml.v2"

	"github.com/spf13/cobra"
)

var initFlags struct {
	//
	ConfigFile string
	//
	Exchange string
	//
	Strategies []struct {
		//
		Market string
		//
		Strategy string
	}
	//
	BTCAddress string
}

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialises the cryptocurrency bot to trade.",
	Long:  ``,
	Run:   initCommand,
}

func init() {
	rootCmd.AddCommand(initCmd)
	initCmd.Flags().StringVar(&initFlags.ConfigFile, "import", "", "imports configuration from file.")
}

func initCommand(cmd *cobra.Command, args []string) {
	if initFlags.ConfigFile != "" {
		c, err := ioutil.ReadFile(initFlags.ConfigFile)
		if err != nil {
			fmt.Println("error opening the configuration file.")
			if GlobalFlags.OutputToScreen > 0 {
				fmt.Printf(": %s", err.Error())
			}
			return
		}

		var flag config.Config

		err = yaml.Unmarshal(c, &flag)

		if err != nil {
			fmt.Println("provided configuration file cannot be loaded.")
			if GlobalFlags.OutputToScreen > 0 {
				fmt.Printf(": %s", err.Error())
			}
			return
		}

		err = ioutil.WriteFile("./bot.yaml", c, 0666)
		if err != nil {
			fmt.Println("cannot create a new configuration file.")
			if GlobalFlags.OutputToScreen > 0 {
				fmt.Printf(": %s", err.Error())
			}
			return
		}
	} else {
		generateConfigFile()
	}
}

// generateConfigFile creates a new configuration file if one has not already been
// created within the bot.
func generateConfigFile() {
	c := config.Config{}

	for {
		var ex config.ExchangeConfig

		fmt.Println("which exchange do you want to include?")
		fmt.Scanln(&ex.ExchangeName)

		exists := false
		for _, exchange := range c.Exchanges {
			if exchange.ExchangeName == ex.ExchangeName {
				exists = true
				break
			}
		}

		if exists {
			fmt.Println("exchange has already been included.")
			continue
		}

		fmt.Println("please provide the Exchange URI...")
		fmt.Scanln(&ex.ExchangeURI)

		fmt.Println("please provide Public Key for the exchange...")
		fmt.Scanln(&ex.PublicKey)

		fmt.Println("please provide Secret Key for the exchange...")
		fmt.Scanln(&ex.SecretKey)

		c.Exchanges = append(c.Exchanges, ex)

		fmt.Println("exchange has been included.")

		var yn string
		for yn != "Y" && yn != "n" {
			fmt.Println("do you want to include another exchange? (Y/n)")
			fmt.Scanln(&yn)
		}

		if yn == "n" {
			break
		}
	}

	for {
		var yn string
		for yn != "Y" && yn != "n" {
			fmt.Println("do you want to add a strategy bound to a market? (Y/n)")
			fmt.Scanln(&yn)
		}

		if yn == "n" {
			break
		}

		var st config.StrategyConfig

		fmt.Println("please enter the name of the strategy you want to use...")
		fmt.Scanln(&st.Strategy)

		for {
			var mc config.MarketConfig
			fmt.Println("please enter market name using short notation e.g. BTC-ETH...")
			fmt.Scanln(&mc.Name)

			for _, exchange := range c.Exchanges {
				var marketName string

				fmt.Printf("please enter %s exchange market ticker or leave empty to skip this exchange.", exchange.ExchangeName)
				fmt.Scanln(&marketName)

				if marketName != "" {
					mc.Exchanges = append(mc.Exchanges, config.ExchangeBindingConfig{
						Name:       exchange.ExchangeName,
						MarketName: marketName,
					})

					fmt.Printf("exchange %s configured with market name %s", exchange.ExchangeName, marketName)
				} else {
					fmt.Printf("exchange %s skipped", exchange.ExchangeName)
				}
			}

			st.Markets = append(st.Markets, mc)

			var yn string
			for yn != "Y" && yn != "n" {
				fmt.Println("do you want to include another market binding to this strategy? (Y/n)")
				fmt.Scanln(&yn)
			}

			if yn == "n" {
				break
			}
		}

		c.Strategies = append(c.Strategies, st)
	}

	cwr, err := yaml.Marshal(c)
	if err != nil {
		fmt.Println("error creating the new configuration file.")
		if GlobalFlags.OutputToScreen > 0 {
			fmt.Printf(": %s", err.Error())
		}
		return
	}

	fmt.Println("the following details:")
	fmt.Println(string(cwr))
	fmt.Println("will be written to ./bot.yaml, is it correct? (Y/n)")

	var yn string
	for yn != "Y" && yn != "n" {
		fmt.Scanln(&yn)
	}

	if yn == "Y" {
		err := ioutil.WriteFile("./bot.yaml", cwr, 0666)
		if err != nil {
			fmt.Println("error while writing to the configuration file.")
			if GlobalFlags.OutputToScreen > 0 {
				fmt.Printf(": %s", err.Error())
			}
		} else {
			fmt.Println("configuration file was created.")
		}
		return
	}

	fmt.Println("You chose not to create a configuration file.")
}
