package generate

import (
	"context"
	"log"
	"os"
	"os/signal"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/weworksandbox/lingo/internal/generator"
	"github.com/weworksandbox/lingo/internal/parse"
)

func Command() *cobra.Command {
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

	ctx, cancel := context.WithCancel(context.Background())

	sigs := make(chan os.Signal)
	signal.Notify(sigs, os.Interrupt, os.Kill)
	go func() {
		<-sigs
		cancel()
	}()

	parser, err := parse.NewMySQL(ctx, s.DataSourceName())
	if err != nil {
		log.Fatalf("unable to connect to database: %s", err)
	}

	if err = generator.Generate(ctx, s, parser); err != nil {
		log.Fatalf("ERR: %s", err)
	}
	log.Println("Completed")
}
