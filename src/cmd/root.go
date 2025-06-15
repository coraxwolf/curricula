package cmd

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/coraxwolf/curricula/src/utils/config"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "curricula",
	Short:   "Curricula is a CLI tool for interacting with the Curricula API",
	Long:    `Curricula is a CLI tool for interacting with the Curricula API`,
	Version: fmt.Sprintf("%s-%s (%s) build-date: %s", config.VERSION, config.BUILD_NUMBER, config.BUILD_STATUS, config.BUILD_DATE),
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.CompletionOptions.DisableDefaultCmd = true

}

func initConfig() {
	// check for configuration file
	ucfg, err := os.UserConfigDir()
	if err != nil {
		slog.Default().Error("Failed to get user config directory", "error", err)
		fmt.Printf("Failed to get user config directory: %v\nPlease use the config_dir flag to specify a directory for your configuration file.\nFlag Values or Defaults will be used.", err)
		return
	}
	configFile := fmt.Sprintf("%s/%s/config.json", ucfg, config.APP_NAME)
	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		// attempt to create the config directory
		if err := os.MkdirAll(fmt.Sprintf("%s/%s", ucfg, config.APP_NAME), 0755); err != nil {
			slog.Default().Error("Failed to create config directory", "error", err)
			fmt.Printf("Failed to create config directory: %v\nPlease use the config_dir flag to specify a directory for your configuration file.\nFlag Values or Defaults will be used.", err)
			return
		}
		// create the config file with default values
		if err := config.WriteConfig(configFile); err != nil {
			slog.Default().Error("Failed to create config file", "error", err)
			fmt.Printf("Failed to create config file: %v\nPlease use the config_dir flag to specify a directory for your configuration file.\nFlag Values or Defaults will be used.", err)
			return
		}
	} else if err != nil {
		slog.Default().Error("Failed to stat config file", "error", err)
		fmt.Printf("Failed to stat config file: %v\nPlease use the config_dir flag to specify a directory for your configuration file.\nFlag Values or Defaults will be used.", err)
		return
	}
	// read the config file
	if err := config.ReadConfig(configFile); err != nil {
		slog.Default().Error("Failed to read config file", "error", err)
		fmt.Printf("Failed to read config file: %v\nPlease use the config_dir flag to specify a directory for your configuration file.\nFlag Values or Defaults will be used.", err)
		return
	}
}

func Execute() error {
	return rootCmd.Execute()
}
