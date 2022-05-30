# ge-dw-admin

![GitHub release (latest by date)](https://img.shields.io/github/v/release/grassrootseconomics/ge-staff-cli)
![GitHub Workflow Status](https://img.shields.io/github/workflow/status/grassrootseconomics/ge-staff-cli/goreleaser)

> GE DW Admin CLI (Admin operations on the Datwarehouse )

## Usage

```bash
NAME:
   ge-dw-admin - GE Back Office Ops CLI

USAGE:
   ge-dw-admin [global options] command [command options] [arguments...]

COMMANDS:
   export-balances  Export balances by token to a CSV file
   help, h          Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --batch-balance value  Batch Balance Smart Contract (default: "0xb9e215B789e9Ec6643Ba4ff7b98EA219F38c6fE5") [$BATCH_BALANCE]
   --dsn value            Warehouse DB DSN [$DB_DSN]
   --help, -h             show help (default: false)
   --rpc value            Kitabu Chain RPC (default: "https://rpc.sarafu.network") [$RPC_PROVIDER]
   --token-index value    Sarafu Network Token Index (default: "0x5A1EB529438D8b3cA943A45a48744f4c73d1f098") [$TOKEN_INDEX]
```
