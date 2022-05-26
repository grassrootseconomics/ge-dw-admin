package main

import (
	"fmt"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/urfave/cli/v2"
)

func init() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
}

func main() {
	app := &cli.App{
		Name:  "export-balance",
		Usage: "Export token balances",
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
				Name:    "batch-balance",
				Value:   "0xb9e215B789e9Ec6643Ba4ff7b98EA219F38c6fE5",
				Usage:   "Batch Balance Smart Contract",
				EnvVars: []string{"BATCH_BALANCE"},
			},
		},
		Commands: []*cli.Command{
			{
				Name:      "token",
				Usage:     "balances by token",
				ArgsUsage: "[token symbol]",
				Action: func(c *cli.Context) error {
					token := c.Args().Get(0)
					fmt.Println(token)
					return nil
				},
			},
			{
				Name:      "user",
				Usage:     "balances by user address",
				ArgsUsage: "[wallet address]",
				Action: func(c *cli.Context) error {
					userAddress := c.Args().Get(0)
					fmt.Println(userAddress)
					return nil
				},
			},
		},
		Before: func(c *cli.Context) error {
			if err := connectDb(c.String("dsn")); err != nil {
				log.Fatal().Err(err).Msg("failed to connect to postgres")
			}

			return nil
		},
		After: func(c *cli.Context) error {
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal().Err(err).Msg("could not run CLI")
	}
}
