package generate

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/weworksandbox/lingo/internal/generator"
	"github.com/weworksandbox/lingo/internal/parse"
)

func Generate() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "generate",
		Short: "Generate entity table and columns from an existing database schema",
		Args:  cobra.MaximumNArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			return generate()
		},
	}

	const (
		flagDir    = "dir"
		flagSchema = "schema"
		flagDriver = "driver"
		flagDSN    = "dsn"
	)
	cmd.PersistentFlags().StringP(flagDir, "d", "./db", "directory where generated file structure should go")
	_ = viper.BindPFlag(flagDir, cmd.Flag(flagDir))
	_ = cmd.PersistentFlags().SetAnnotation(flagDir, cobra.BashCompSubdirsInDir, []string{})

	cmd.PersistentFlags().StringSliceP(flagSchema, "s", []string{}, "schema name to generate for")
	_ = viper.BindPFlag(flagSchema, cmd.Flag(flagSchema))

	cmd.PersistentFlags().String(flagDriver, "mysql", "driver name used to initialize the SQL driver")
	_ = viper.BindPFlag(flagDriver, cmd.Flag(flagDriver))

	cmd.PersistentFlags().String(flagDSN, "", "data source connection string")
	_ = viper.BindPFlag(flagDSN, cmd.Flag(flagDSN))
	return cmd
}

func generate() error {
	var settings = getSettings()

	parser, err := parse.NewMySQL(settings.DataSourceName())
	if err != nil {
		return err
	}

	var combined error
	for err := range generator.Generate(context.Background(), settings, parser) {
		if combined == nil {
			combined = err
		} else {
			combined = fmt.Errorf("%s: %w", combined, err)
		}
	}
	return combined
}
