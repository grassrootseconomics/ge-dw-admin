package main

import (
	"context"
	"encoding/csv"
	"fmt"
	"os"
	"time"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/lmittmann/w3"
	"github.com/rs/zerolog/log"
)

type tokenHoldersRes struct {
	Address string `db:"holding_address"`
	Phone   string `db:"phone_number"`
}

func exportBalancestoCSV(ctx context.Context, tokenSymbol string) error {
	var res []tokenHoldersRes
	var tokenAddressRaw string

	if err := db.QueryRow(ctx, tokenAddressFromSymbol, tokenSymbol).Scan(&tokenAddressRaw); err != nil {
		return err
	}
	tokenAddressChecksumed := w3.A(checksum(tokenAddressRaw))
	log.Info().Msgf("%s address is %s", tokenSymbol, tokenAddressChecksumed)

	rows, err := db.Query(ctx, tokenHolders, tokenAddressRaw)
	if err != nil {
		return err
	}

	if err := pgxscan.ScanAll(&res, rows); err != nil {
		return err
	}

	log.Info().Msgf("%d %s holders", len(res), tokenSymbol)
	log.Info().Msgf("starting to dump all balances...")

	fName := fmt.Sprintf("%s_balances_%s.csv", tokenSymbol, time.Now().Format("20060102150405"))
	log.Info().Msgf("creating file %s", fName)
	f, err := os.Create(fName)
	if err != nil {
		return err
	}

	defer f.Close()

	w := csv.NewWriter(f)
	if err := w.Error(); err != nil {
		return err
	}
	defer w.Flush()

	headers := []string{
		"voucher_symbol",
		"holder_address",
		"phone",
		"balance",
	}

	w.Write(headers)

	for i, r := range res {
		row := make([]string, 0, 1+len(headers))

		if i%100 == 0 || i == len(res)-1 {
			log.Info().Msgf("dumped %d balances", i+1)
		}

		checksumedAddress := w3.A(checksum(r.Address))
		balance, err := cicnetClient.BalanceOf(ctx, tokenAddressChecksumed, checksumedAddress)
		if err != nil {
			return err
		}

		row = append(
			row,
			tokenSymbol,
			r.Address,
			r.Phone,
			balance.String(),
		)

		w.Write(row)
	}

	return nil
}
