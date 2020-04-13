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

	cmd.PersistentFlags().StringP("dir", "d", "./db", "directory where generated file structure should go")
	_ = viper.BindPFlag("dir", cmd.Flag("dir"))
	_ = cmd.PersistentFlags().SetAnnotation("dir", cobra.BashCompSubdirsInDir, []string{})

	cmd.PersistentFlags().StringSliceP("schema", "s", []string{}, "schema name to generate for")
	_ = viper.BindPFlag("schema", cmd.Flag("schema"))

	cmd.PersistentFlags().String("driver", "mysql", "driver name used to initialize the SQL driver")
	_ = viper.BindPFlag("driver", cmd.Flag("driver"))

	cmd.PersistentFlags().String("dsn", "", "data source connection string")
	_ = viper.BindPFlag("dsn", cmd.Flag("dsn"))
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
