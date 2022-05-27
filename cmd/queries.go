package main

const (
	tokenHolders = `WITH unique_holders AS (
		SELECT sender_address AS holding_address FROM transactions
		  WHERE token_address = $1
		  UNION
		  SELECT recipient_address AS holding_address FROM transactions
		  WHERE token_address = $1
	)
	
	SELECT holding_address, users.phone_number FROM unique_holders
	INNER JOIN users ON unique_holders.holding_address = users.blockchain_address;`

	tokenAddressFromSymbol = "SELECT token_address FROM tokens WHERE token_symbol = $1"
)
