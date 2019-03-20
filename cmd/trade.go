package cfx

import (
	"fmt"

	"github.com/spf13/cobra"
)

var tradeCmd = &cobra.Command{
	Use:   "trade",
	Short: "",
	Long:  ``,
	Run:   tradeCommand,
}

var tradeFlags struct {
	Simulate bool
}

func init() {
	rootCmd.AddCommand(tradeCmd)
	tradeCmd.Flags().BoolVarP(&tradeFlags.Simulate, "simulate", "s", false, "simulate the trades instead of actually running them.")
}

func tradeCommand(cmd *cobra.Command, args []string) {
	fmt.Print("Trade yo!")
}
