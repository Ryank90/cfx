package cfx

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/spf13/cobra"
)

const (
	version = "0.0.1-alpha"
)

var signals chan os.Signal

var rootFlags struct {
	Version bool
}

var rootCmd = &cobra.Command{
	Use:   "cfx",
	Short: "CFX cryptocurrency trading bot.",
	Long:  "CFX is a cryptocurrency trading bot that interfaces with multiple exchanges to perform buy and sell actions on your behalf.",
	Run: func(cmd *cobra.Command, args []string) {
		if rootFlags.Version {
			fmt.Printf("%s\n", version)
		} else {
			cmd.Help()
		}
	},
}

func init() {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	go func() {
		<-signals
		signal.Stop(signals)
		fmt.Println("CTRL-C command recieved... Exiting...")
		os.Exit(0)
	}()

	rootCmd.Flags().BoolVarP(&rootFlags.Version, "version", "v", false, "show version information.")
	rootCmd.PersistentFlags().StringVar(&GlobalFlags.ConfigFile, "config", "bot.yaml", "configuration file path (default: ./bot.yaml).")
	rootCmd.PersistentFlags().CountVarP(&GlobalFlags.OutputToScreen, "output", "o", "show verbose information when trading.")
}

// Execute is the entry point for the CLI.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
