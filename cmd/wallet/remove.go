package wallet

import (
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler/wallet"
	"github.com/urfave/cli/v2"
)

var RemoveCmd = &cli.Command{
	Name:      "remove",
	Usage:     "Remove a wallet",
	ArgsUsage: "<address>",
	Action: func(c *cli.Context) error {
		db := database.MustOpenFromCLI(c)
		err := wallet.RemoveHandler(db, c.Args().Get(0))

		if err != nil {
			return err.CliError()
		}

		return nil
	},
}
