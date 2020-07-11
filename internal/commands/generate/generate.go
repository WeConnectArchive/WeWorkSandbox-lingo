package generate

import (
	"context"
	"log"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/weworksandbox/lingo/internal/generator"
	"github.com/weworksandbox/lingo/internal/parse"
)

func Generate() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "generate",
		Short: "Generate tables from an existing database schema",
		Args:  cobra.MaximumNArgs(0),
		Run:   generate,
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

//revive:disable:deep-exit - Disabling deep exits from log.Fatalf due to this being a 'top level' command.

func generate(_ *cobra.Command, _ []string) {
	var s = getSettings()

	switch dn := strings.ToLower(s.DriverName()); dn {
	case "mysql":
		// TODO refactor MySQL Parser into Interface
	default:
		log.Fatalf("parser unknown for driver '%s'", dn)
	}

	parser, err := parse.NewMySQL(s.DataSourceName())
	if err != nil {
		log.Fatalf("unable to connect to database: %s", err)
	}

	var errCount int
	for err = range generator.Generate(context.Background(), s, parser) {
		log.Printf("ERR: %s", err)
		errCount++
	}
	if errCount != 0 {
		log.Fatalf("had %d errors occur while generating", errCount)
	}
}
