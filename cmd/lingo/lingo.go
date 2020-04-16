package main

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/weworksandbox/lingo/internal/commands/generate"
)

func main() {
	var configFile string

	// Run `initConfig` before each command below is executed
	cobra.OnInitialize(func() {
		viper.SetConfigFile(configFile)
		if err := viper.ReadInConfig(); err != nil {
			log.Fatalf("Can't find or read config at '%s': %s", configFile, err)
		}
	})

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
