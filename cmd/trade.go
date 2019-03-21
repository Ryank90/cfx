package cfx

import (
	"cfx/config"
	"cfx/exchange"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/spf13/cobra"
	yaml "gopkg.in/yaml.v2"
)

var tradeCmd = &cobra.Command{
	Use:   "trade",
	Short: "Starts the bot to perform trades from saved configuration.",
	Long:  ``,
	Run:   tradeCommand,
}

var tradeFlags struct {
	Simulate bool
}

var cfg config.Config

func init() {
	rootCmd.AddCommand(tradeCmd)
	tradeCmd.Flags().BoolVarP(&tradeFlags.Simulate, "simulate", "s", false, "simulate the trades instead of actually running them.")
}

func initConfig() error {
	c, err := os.Open(GlobalFlags.ConfigFile)
	if err != nil {
		return err
	}

	cm, err := ioutil.ReadAll(c)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(cm, &cfg)
	if err != nil {
		return err
	}

	return nil
}

func tradeCommand(cmd *cobra.Command, args []string) {
	fmt.Println("Getting bot configuration...")
	if err := initConfig(); err != nil {
		fmt.Println("Cannot read the current bot.yaml file. Please create or replace it by running `init` command.")
		return
	}
	fmt.Println("Completed!")

	fmt.Println("Getting exchange information...")
	w := make([]exchange.Exchange, len(cfg.Exchanges))
	for i, excfg := range cfg.Exchanges {
		w[i] = exchange.InitExchange(excfg, cfg.SimulationMode, excfg.TestBalances, excfg.DepositAddress)
	}
	fmt.Println("Completed!")
}
