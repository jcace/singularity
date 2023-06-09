package dataset

import (
	"github.com/data-preservation-programs/singularity/cmd/cliutil"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler/dataset"
	"github.com/urfave/cli/v2"
)

var ListPiecesCmd = &cli.Command{
	Name:      "list-pieces",
	Usage:     "[alpha] List all pieces for the dataset that are available for deal making",
	ArgsUsage: "<dataset_name>",
	Action: func(c *cli.Context) error {
		db := database.MustOpenFromCLI(c)

		car, err := dataset.ListPiecesHandler(
			db, c.Args().Get(0),
		)
		if err != nil {
			return err.CliError()
		}

		cliutil.PrintToConsole(car, c.Bool("json"))
		return nil
	},
}
