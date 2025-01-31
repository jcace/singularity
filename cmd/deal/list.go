package deal

import (
	"github.com/data-preservation-programs/singularity/cmd/cliutil"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler/deal"
	"github.com/urfave/cli/v2"
)

var ListCmd = &cli.Command{
	Name:  "list",
	Usage: "List all deals",
	Flags: []cli.Flag{
		&cli.StringSliceFlag{
			Name: "dataset",
			Usage: "Filter deals by dataset name",
		},
		&cli.UintSliceFlag{
			Name: "schedule",
			Usage: "Filter deals by schedule",
		},
		&cli.StringSliceFlag{
			Name: "provider",
			Usage: "Filter deals by provider",
		},
		&cli.StringSliceFlag{
			Name: "state",
			Usage: "Filter deals by state: proposed, published, active, expired, proposal_expired, slashed",
		},
	},
	Action: func(c *cli.Context) error {
		db := database.MustOpenFromCLI(c)
		deals, err := deal.ListHandler(db, deal.ListDealRequest{

		})
		if err != nil {
			return err.CliError()
		}
		cliutil.PrintToConsole(deals, c.Bool("json"))
		return nil
	},
}
