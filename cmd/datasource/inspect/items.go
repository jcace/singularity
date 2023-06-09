package inspect

import (
	"github.com/data-preservation-programs/singularity/cmd/cliutil"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler/datasource/inspect"
	"github.com/urfave/cli/v2"
)

var ItemsCmd = &cli.Command{
	Name:        "items",
	Usage:       "Get all item details of a data source",
	ArgsUsage:   "<source_id>",
	Description: "This command will list all items in a data source. This may be very large list.",
	Action: func(c *cli.Context) error {
		db := database.MustOpenFromCLI(c)
		result, err := inspect.GetSourceItemsHandler(
			db,
			c.Args().Get(0),
		)
		if err != nil {
			return err.CliError()
		}

		cliutil.PrintToConsole(result, c.Bool("json"))
		return nil
	},
}
