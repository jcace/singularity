package ez

import (
	"fmt"
	"github.com/data-preservation-programs/singularity/cmd/cliutil"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler/admin"
	"github.com/data-preservation-programs/singularity/handler/dataset"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/service/datasetworker"
	"github.com/pkg/errors"
	"github.com/urfave/cli/v2"
	"gorm.io/gorm"
	"os"
	"path/filepath"
	"strings"
	"time"
)

var PrepCmd = &cli.Command{
	Name:      "ez-prep",
	Category:  "Easy Commands",
	ArgsUsage: "<path>",
	Usage:     "Prepare a dataset from a local path",
	Description: "This commands can be used to prepare a dataset from a local path with minimum configurable parameters.\n" +
		"For more advanced usage, please use the subcommands under `dataset` and `datasource`.\n" +
		"You can also use this command for benchmarking with in-memory database and inline preparation, i.e.\n" +
		"  mkdir dataset\n" +
		"  truncate -s 1024G test.img\n" +
		"  singularity ez-prep --output-dir '' --database-file '' -j $(nproc) ./dataset",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    "max-size",
			Aliases: []string{"M"},
			Usage:   "Maximum size of the CAR files to be created",
			Value:   "31.5GiB",
		},
		&cli.StringFlag{
			Name:    "output-dir",
			Aliases: []string{"o"},
			Usage:   "Output directory for CAR files. To use inline preparation, use an empty string",
			Value:   "./cars",
		},
		&cli.IntFlag{
			Name:    "concurrency",
			Aliases: []string{"j"},
			Usage:   "Concurrency for packing",
			Value:   1,
		},
		&cli.StringFlag{
			Name:        "database-file",
			Usage:       "The database file to store the metadata. To use in memory database, use an empty string.",
			DefaultText: "./ezprep-<name>.db",
		},
	},
	Action: func(c *cli.Context) error {
		t := time.Now().Unix()
		path := c.Args().Get(0)
		if path == "" {
			return errors.New("path is required")
		}
		databaseFile := c.String("database-file")
		if databaseFile == "" {
			if c.IsSet("database-file") {
				databaseFile = "file::memory:?cache=shared"
			} else {
				databaseFile = fmt.Sprintf("./ezprep-%d.db", t)
			}
		}
		var err error
		if !strings.HasPrefix(databaseFile, "file::memory") {
			databaseFile, err = filepath.Abs(databaseFile)
			if err != nil {
				return errors.Wrap(err, "failed to get absolute path")
			}
		}
		db, err := database.Open("sqlite:"+databaseFile, &gorm.Config{})
		if err != nil {
			return errors.Wrapf(err, "failed to open database %s", databaseFile)
		}

		// Step 1, initialize the database
		err = admin.InitHandler(db).CliError()
		if err != nil {
			return err
		}

		// Step 2, create a dataset
		var outputDirs []string
		if c.String("output-dir") != "" {
			outputDirs = []string{c.String("output-dir")}
			err = os.MkdirAll(outputDirs[0], 0755)
			if err != nil {
				return errors.Wrap(err, "failed to create output directory")
			}
		}
		ds, err2 := dataset.CreateHandler(db, dataset.CreateRequest{
			Name:       "ez",
			MaxSizeStr: c.String("max-size"),
			OutputDirs: outputDirs,
		})
		if err2 != nil {
			return err2.CliError()
		}

		// Step 3, add a local data source
		path, err = filepath.Abs(path)
		if err != nil {
			return errors.Wrap(err, "failed to get absolute path")
		}
		root := model.Directory{}
		err = db.Create(&root).Error
		if err != nil {
			return errors.Wrap(err, "failed to create root directory")
		}
		source := model.Source{
			DatasetID:       ds.ID,
			Type:            "local",
			Path:            path,
			Metadata:        model.Metadata(nil),
			ScanningState:   model.Ready,
			RootDirectoryID: root.ID,
		}
		err = db.Create(&source).Error
		if err != nil {
			return errors.Wrap(err, "failed to create source")
		}

		// Step 3, start dataset worker
		worker := datasetworker.NewDatasetWorker(
			db,
			datasetworker.DatasetWorkerConfig{
				Concurrency:    c.Int("concurrency"),
				EnableScan:     true,
				EnablePack:     true,
				EnableDag:      true,
				ExitOnComplete: true,
			})
		err = worker.Run(c.Context)
		if err != nil {
			return err
		}

		// Step 4, print all information
		cars, err2 := dataset.ListPiecesHandler(
			db, ds.Name,
		)
		if err2 != nil {
			return err2.CliError()
		}

		cliutil.PrintToConsole(cars, false)
		return nil
	},
}
