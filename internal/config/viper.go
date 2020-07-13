package config

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

// ConfigFile is the location of the configuration file.
var File string

// FileFlag is used for initializing a configuration flag
var FileFlag *pflag.FlagSet

var configFileBashCompletion = []string{"!*.yml", "!*.yaml", "!*.properties", "!*.toml", "!*.json"}

func init() {
	FileFlag = pflag.NewFlagSet("configuration", pflag.ContinueOnError)
	FileFlag.StringVar(&File, "config", "", "configuration file")
	_ = FileFlag.SetAnnotation("config", cobra.BashCompFilenameExt, configFileBashCompletion)

	cobra.OnInitialize(func() {
		err := ReadConfig()
		if err != nil {
			log.Fatal(err) // Print the error if there is one
		}
	})
}

func ReadConfig() error {
	viper.SetConfigFile(File)

	err := viper.ReadInConfig()
	if err != nil {
		// Replace this with errors.Is once this conforms to 1.13 errors.
		if _, notFound := err.(viper.ConfigFileNotFoundError); !notFound {
			return fmt.Errorf("unable to read config at '%s': %w", File, err)
		}
	}
	return nil
}
