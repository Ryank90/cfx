package cfx

import (
	"fmt"

	"github.com/spf13/cobra"
)

var initFlags struct {
	ConfigFile string
	Exchange   string
	Strategies []struct {
		Market   string
		Strategy string
	}
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
	fmt.Print("Init yo!")
}
