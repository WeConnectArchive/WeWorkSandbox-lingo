package main

import (
	"github.com/spf13/cobra"

	"github.com/weworksandbox/lingo/internal/commands/generate"
	"github.com/weworksandbox/lingo/internal/config"
)

func main() {
	var rootCmd = &cobra.Command{
		Use: "lingo",
	}

	rootCmd.PersistentFlags().AddFlagSet(config.FileFlag)
	rootCmd.AddCommand(generate.Generate())
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}
