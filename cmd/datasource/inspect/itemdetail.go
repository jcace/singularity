package inspect

import (
	"fmt"
	"github.com/data-preservation-programs/singularity/cmd/cliutil"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler/datasource/inspect"
	"github.com/urfave/cli/v2"
)

var ItemDetailCmd = &cli.Command{
	Name:      "itemdetail",
	Usage:     "Get details about a specific item",
	ArgsUsage: "<item_id>",
	Action: func(c *cli.Context) error {
		db := database.MustOpenFromCLI(c)
		result, err := inspect.GetSourceItemDetailHandler(
			db,
			c.Args().Get(0),
		)
		if err != nil {
			return err.CliError()
		}

		if c.Bool("json") {
			cliutil.PrintToConsole(result, true)
			return nil
		}

		fmt.Println("Item:")
		cliutil.PrintToConsole(result, false)
		fmt.Println("Item Parts:")
		cliutil.PrintToConsole(result.ItemParts, false)
		return nil
	},
}
