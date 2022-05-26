package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "export-balance",
		Usage: "Export token balances",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "host",
				Usage:    "Warehouse DB Host",
				Required: true,
				EnvVars:  []string{"DB_HOST"},
			},
			&cli.StringFlag{
				Name:     "port",
				Usage:    "Warehouse DB Port",
				Required: true,
				EnvVars:  []string{"DB_PORT"},
			},
			&cli.StringFlag{
				Name:     "username",
				Usage:    "Warehouse DB Username",
				Required: true,
				EnvVars:  []string{"DB_USERNAME"},
			},
			&cli.StringFlag{
				Name:     "password",
				Usage:    "Warehouse DB Password",
				Required: true,
				EnvVars:  []string{"DB_PASSWORD"},
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
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
