package cfx

import (
	"cfx/config"
	"fmt"
	"io/ioutil"

	"github.com/spf13/cobra"
	yaml "gopkg.in/yaml.v2"
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
	fmt.Println("generate a file!")
}
