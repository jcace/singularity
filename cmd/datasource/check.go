package datasource

import (
	"github.com/data-preservation-programs/singularity/cmd/cliutil"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler/datasource"
	"github.com/urfave/cli/v2"
)

var CheckCmd = &cli.Command{
	Name:      "check",
	Usage:     "Check a data source by listing its entries. This is not to list prepared items but a direct listing of the data source to verify datasource connection",
	ArgsUsage: "<source_id> [sub_path]",
	Description: "This command will list entries in a data source under <sub_path>. " +
		"If <sub_path> is not provided, it will use the root directory",
	Action: func(c *cli.Context) error {
		db := database.MustOpenFromCLI(c)
		entries, err := datasource.CheckSourceHandler(
			db,
			c.Context,
			c.Args().Get(0),
			datasource.CheckSourceRequest{
				Path: c.Args().Get(1),
			},
		)
		if err != nil {
			return err.CliError()
		}

		cliutil.PrintToConsole(entries, c.Bool("json"))
		return nil
	},
}
