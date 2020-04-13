package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/weworksandbox/lingo/internal/commands/generate"
)

var configFile string

func initConfig() {
	if configFile != "" {
		viper.SetConfigFile(configFile)
		if err := viper.ReadInConfig(); err != nil {
			fmt.Printf("Can't find or read config at '%s': %s", configFile, err)
			os.Exit(1)
		}
	}
}

func main() {
	// Run `initConfig` before each command below is executed
	cobra.OnInitialize(initConfig)

	var rootCmd = &cobra.Command{
		Use: "lingo",
	}

	rootCmd.PersistentFlags().StringVar(&configFile, "config", "", "configuration file")
	_ = rootCmd.PersistentFlags().SetAnnotation("config", cobra.BashCompFilenameExt, []string{"!*.yml", "!*.yaml", "!*.properties", "!*.toml", "!*.json"})
	rootCmd.AddCommand(generate.Generate())
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}
