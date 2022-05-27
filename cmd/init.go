package main

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	cic_net "github.com/grassrootseconomics/cic-go/net"
	"github.com/grassrootseconomics/cic-go/provider"
	"github.com/jackc/pgx/v4/pgxpool"
)

var (
	db           *pgxpool.Pool
	rpcProvider  *provider.Provider
	cicnetClient *cic_net.CicNet
)

func connectDb(dsn string) error {
	var err error
	db, err = pgxpool.Connect(context.Background(), dsn)
	if err != nil {
		return err
	}

	return nil
}

func loadProvider(rpcEndpoint string) error {
	var err error

	rpcProvider, err = provider.NewRpcProvider(rpcEndpoint)
	if err != nil {
		return err
	}

	return nil
}

func loadCicNet(tokenIndex common.Address) error {
	var err error

	cicnetClient, err = cic_net.NewCicNet(rpcProvider, tokenIndex)
	if err != nil {
		return err
	}

	return nil
}
