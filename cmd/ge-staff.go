package main

import (
	"os"

	"github.com/lmittmann/w3"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/urfave/cli/v2"
)

func init() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
}

func main() {
	app := &cli.App{
		Name:  "ge-staff",
		Usage: "GE Back Office Ops CLI",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "dsn",
				Usage:    "Warehouse DB DSN",
				Required: true,
				EnvVars:  []string{"DB_DSN"},
			},
			&cli.StringFlag{
				Name:    "rpc",
				Value:   "https://rpc.sarafu.network",
				Usage:   "Kitabu Chain RPC",
				EnvVars: []string{"RPC_PROVIDER"},
			},
			&cli.StringFlag{
				Name:    "token-index",
				Value:   "0x5A1EB529438D8b3cA943A45a48744f4c73d1f098",
				Usage:   "Sarafu Network Token Index",
				EnvVars: []string{"TOKEN_INDEX"},
			},
			&cli.StringFlag{
				Name:    "batch-balance",
				Value:   "0xb9e215B789e9Ec6643Ba4ff7b98EA219F38c6fE5",
				Usage:   "Batch Balance Smart Contract",
				EnvVars: []string{"BATCH_BALANCE"},
			},
		},
		Commands: []*cli.Command{
			{
				Name:      "export-balances",
				Usage:     "Export balances by token to a CSV file",
				ArgsUsage: "[token symbol]",
				Action: func(c *cli.Context) error {
					token := c.Args().Get(0)
					if err := exportBalancestoCSV(c.Context, token); err != nil {
						log.Fatal().Err(err).Msgf("failed %s", err)
						return err
					}
					return nil
				},
			},
		},
		Before: func(c *cli.Context) error {
			if err := connectDb(c.String("dsn")); err != nil {
				log.Fatal().Err(err).Msg("failed to connect to dw")
			}

			if err := loadProvider(c.String("rpc")); err != nil {
				log.Fatal().Err(err).Msg("failed to connect to rpc endpoint")
			}

			if err := loadCicNet(w3.A(c.String("token-index"))); err != nil {
				log.Fatal().Err(err).Msg("failed to load cicnet")
			}
			log.Info().Msg("successfully conected to data warehouse")

			return nil
		},
		After: func(c *cli.Context) error {
			db.Close()
			log.Info().Msg("closed all data warehouse connections")
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal().Err(err).Msg("could not run ge-staff CLI")
	}
}
