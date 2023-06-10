package cmd

import (
	"context"
	"fmt"
	"github.com/data-preservation-programs/singularity/cmd/admin"
	"github.com/data-preservation-programs/singularity/cmd/dataset"
	"github.com/data-preservation-programs/singularity/cmd/datasource"
	"github.com/data-preservation-programs/singularity/cmd/datasource/inspect"
	"github.com/data-preservation-programs/singularity/cmd/deal"
	"github.com/data-preservation-programs/singularity/cmd/deal/schedule"
	"github.com/data-preservation-programs/singularity/cmd/deal/spadepolicy"
	"github.com/data-preservation-programs/singularity/cmd/ez"
	"github.com/data-preservation-programs/singularity/cmd/run"
	"github.com/data-preservation-programs/singularity/cmd/wallet"
	"github.com/data-preservation-programs/singularity/util/must"
	"github.com/mattn/go-shellwords"
	"github.com/urfave/cli/v2"
	"io"
	"os"
)

var app = &cli.App{
	Name:                 "singularity",
	Usage:                "A tool for large-scale clients with PB-scale data onboarding to Filecoin network",
	EnableBashCompletion: true,
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name: "database-connection-string",
			Usage: "Connection string to the database.\n" +
				"Supported database: sqlite3, postgres, mysql, sqlserver\n" +
				"Example for postgres  - postgres://user:pass@example.com:5432/dbname\n" +
				"Example for mysql     - mysql://user:pass@tcp(localhost:3306)/dbname?charset=ascii&parseTime=true\n" +
				"                          Note: the database needs to be created using ascii Character Set:" +
				"                                `CREATE DATABASE <dbname> DEFAULT CHARACTER SET ascii`\n" +
				"Example for sqlserver - sqlserver://user:pass@example.com:9930?database=dbname\"\n" +
				"Example for sqlite3   - sqlite:/absolute/path/to/database.db\n" +
				"            or        - sqlite:relative/path/to/database.db\n",
			DefaultText: "sqlite:" + must.String(os.UserHomeDir()) + "/.singularity/singularity.db",
			Value:       "sqlite:" + must.String(os.UserHomeDir()) + "/.singularity/singularity.db",
			EnvVars:     []string{"DATABASE_CONNECTION_STRING"},
		},
		&cli.BoolFlag{
			Name:  "verbose",
			Usage: "Enable verbose logging",
			Value: false,
		},
		&cli.BoolFlag{
			Name:  "json",
			Usage: "Enable JSON output",
			Value: false,
		},
	},
	Commands: []*cli.Command{
		ez.PrepCmd,
		{
			Name:     "admin",
			Usage:    "Admin commands",
			Category: "Operations",
			Subcommands: []*cli.Command{
				admin.InitCmd,
				admin.ResetCmd,
				admin.MigrateCmd,
			},
		},
		DownloadCmd,
		{
			Name:     "deal",
			Usage:    "Replication / Deal making management",
			Category: "Operations",
			Subcommands: []*cli.Command{
				{
					Name:  "schedule",
					Usage: "Schedule deals",
					Subcommands: []*cli.Command{
						schedule.CreateCmd,
						schedule.ListCmd,
						schedule.PauseCmd,
						schedule.ResumeCmd,
					},
				},
				{
					Name:  "spade-policy",
					Usage: "Manage SPADE policies",
					Subcommands: []*cli.Command{
						spadepolicy.CreateCmd,
						spadepolicy.ListCmd,
						spadepolicy.RemoveCmd,
					},
				},
				deal.SendManualCmd,
				deal.ListCmd,
			},
		},
		{
			Name:     "run",
			Category: "Daemons",
			Usage:    "Run different singularity components",
			Subcommands: []*cli.Command{
				run.ApiCmd,
				run.DatasetWorkerCmd,
				run.ContentProviderCmd,
				run.DealMakerCmd,
				run.SpadeAPICmd,
			},
		},
		{
			Name:     "dataset",
			Category: "Operations",
			Usage:    "Dataset management",
			Subcommands: []*cli.Command{
				dataset.CreateCmd,
				dataset.ListDatasetCmd,
				dataset.UpdateCmd,
				dataset.RemoveDatasetCmd,
				dataset.AddWalletCmd,
				dataset.ListWalletCmd,
				dataset.RemoveWalletCmd,
				dataset.AddPieceCmd,
				dataset.ListPiecesCmd,
			},
		},
		{
			Name:     "datasource",
			Category: "Operations",
			Usage:    "Data source management",
			Subcommands: []*cli.Command{
				datasource.AddCmd,
				datasource.ListCmd,
				datasource.StatusCmd,
				datasource.RemoveCmd,
				datasource.CheckCmd,
				datasource.UpdateCmd,
				datasource.RescanCmd,
				{
					Name:  "inspect",
					Usage: "Get preparation status of a data source",
					Subcommands: []*cli.Command{
						inspect.ChunksCmd,
						inspect.ItemsCmd,
						inspect.ChunkDetailCmd,
						inspect.ItemDetailCmd,
						inspect.DirCmd,
					},
				},
			},
		},
		{
			Name:     "wallet",
			Category: "Operations",
			Usage:    "Wallet management",
			Subcommands: []*cli.Command{
				wallet.ImportCmd,
				wallet.ListCmd,
				wallet.AddRemoteCmd,
				wallet.RemoveCmd,
			},
		},
	},
}

func RunApp(ctx context.Context, args []string) error {
	if err := app.RunContext(ctx, args); err != nil {
		return err
	}

	return nil
}

func RunArgsInTest(ctx context.Context, args string) (string, string, error) {
	app.ExitErrHandler = func(c *cli.Context, err error) {
		return
	}
	parser := shellwords.NewParser()
	parser.ParseEnv = true // Enable environment variable parsing
	parsedArgs, err := parser.Parse(args)
	if err != nil {
		return "", "", err
	}

	// Create pipes
	rOut, wOut, _ := os.Pipe()
	rErr, wErr, _ := os.Pipe()

	// Save current stdout and stderr
	oldOut := os.Stdout
	oldErr := os.Stderr

	// Overwrite the stdout and stderr
	os.Stdout = wOut
	os.Stderr = wErr

	outC := make(chan string) // Buffered to prevent goroutine leak
	errC := make(chan string)
	go func() {
		out, _ := io.ReadAll(rOut)
		outC <- string(out)
	}()
	go func() {
		out, _ := io.ReadAll(rErr)
		errC <- string(out)
	}()

	err = RunApp(ctx, parsedArgs)

	// Close the pipes
	wOut.Close()
	wErr.Close()

	// Restore original stdout and stderr
	os.Stdout = oldOut
	os.Stderr = oldErr

	// Wait for the output from the goroutines
	outputOut := <-outC
	outputErr := <-errC

	// Let's still print it to stdout and stderr
	fmt.Println(outputOut)
	fmt.Fprintln(os.Stderr, outputErr)
	return outputOut, outputErr, err
}
