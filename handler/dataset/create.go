package dataset

import (
	"github.com/data-preservation-programs/singularity/database"
	"github.com/pkg/errors"
	"os"
	"path/filepath"

	"github.com/data-preservation-programs/singularity/handler"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/util"
	"github.com/dustin/go-humanize"
	"github.com/ipfs/go-log/v2"
	"gorm.io/gorm"
)

type CreateRequest struct {
	Name                 string   `json:"name" validate:"required"`                      // Name must be a unique identifier for a dataset
	MaxSizeStr           string   `json:"maxSize" default:"31.5GiB" validate:"required"` // Maximum size of the CAR files to be created
	PieceSizeStr         string   `json:"pieceSize" default:"" validate:"optional"`      // Target piece size of the CAR files used for piece commitment calculation
	OutputDirs           []string `json:"outputDirs" validate:"optional"`                // Output directory for CAR files. Do not set if using inline preparation
	EncryptionRecipients []string `json:"encryptionRecipients" validate:"optional"`      // Public key of the encryption recipient
	EncryptionScript     string   `json:"encryptionScript" validate:"optional"`          // EncryptionScript command to run for custom encryption
}

func parseCreateRequest(request CreateRequest) (*model.Dataset, *handler.Error) {
	maxSize, err := humanize.ParseBytes(request.MaxSizeStr)
	if err != nil {
		return nil, handler.NewBadRequestString("invalid value for max-size: " + err.Error())
	}

	pieceSize := util.NextPowerOfTwo(maxSize)
	if request.PieceSizeStr != "" {
		pieceSize, err = humanize.ParseBytes(request.PieceSizeStr)
		if err != nil {
			return nil, handler.NewBadRequestString("invalid value for piece-size: " + err.Error())
		}

		if pieceSize != util.NextPowerOfTwo(pieceSize) {
			return nil, handler.NewBadRequestString("piece size must be a power of two")
		}
	}

	if pieceSize > 1<<36 {
		return nil, handler.NewBadRequestString("piece size cannot be larger than 64 GiB")
	}

	if maxSize*128/127 >= pieceSize {
		return nil, handler.NewBadRequestString("max size needs to be reduced to leave space for padding")
	}

	outDirs := make([]string, len(request.OutputDirs))
	for i, outputDir := range request.OutputDirs {
		info, err := os.Stat(outputDir)
		if err != nil || !info.IsDir() {
			return nil, handler.NewBadRequestString("output directory does not exist: " + outputDir)
		}
		abs, err := filepath.Abs(outputDir)
		if err != nil {
			return nil, handler.NewBadRequestString("could not get absolute path for output directory: " + err.Error())
		}
		outDirs[i] = abs
	}

	if len(request.EncryptionRecipients) > 0 && request.EncryptionScript != "" {
		return nil, handler.NewBadRequestString("encryption recipients and script cannot be used together")
	}

	if (len(request.EncryptionRecipients) > 0 || request.EncryptionScript != "") && len(request.OutputDirs) == 0 {
		return nil, handler.NewBadRequestString(
			"encryption is not compatible with inline preparation and " +
				"requires at least one output directory",
		)
	}

	return &model.Dataset{
		Name:                 request.Name,
		MaxSize:              int64(maxSize),
		PieceSize:            int64(pieceSize),
		OutputDirs:           outDirs,
		EncryptionRecipients: request.EncryptionRecipients,
		EncryptionScript:     request.EncryptionScript,
	}, nil
}

// CreateHandler godoc
// @Summary Create a new dataset
// @Tags Dataset
// @Accept json
// @Produce json
// @Description The dataset is a top level object to distinguish different dataset.
// @Param request body CreateRequest true "Request body"
// @Success 200 {object} model.Dataset
// @Failure 400 {object} handler.HTTPError
// @Failure 500 {object} handler.HTTPError
// @Router /dataset [post]
func CreateHandler(
	db *gorm.DB,
	request CreateRequest,
) (*model.Dataset, *handler.Error) {
	logger := log.Logger("cli")
	if request.Name == "" {
		return nil, handler.NewBadRequestString("name is required")
	}

	dataset, err := parseCreateRequest(request)
	if err != nil {
		return nil, err
	}

	err2 := database.DoRetry(func() error { return db.Create(dataset).Error })
	if errors.Is(err2, gorm.ErrDuplicatedKey) {
		return nil, handler.NewBadRequestString("dataset with this name already exists")
	}

	if err2 != nil {
		return nil, handler.NewHandlerError(err2)
	}

	logger.Infof("Dataset created with ID: %d", dataset.ID)
	return dataset, nil
}
