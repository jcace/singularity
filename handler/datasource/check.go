package datasource

import (
	"context"
	"github.com/data-preservation-programs/singularity/datasource"
	"github.com/data-preservation-programs/singularity/handler"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/ipfs/go-log/v2"
	"github.com/pkg/errors"
	"github.com/rclone/rclone/fs"
	"github.com/rjNemo/underscore"
	"gorm.io/gorm"
	"strconv"
)

type CheckSourceRequest struct {
	Path string `json:"path" validate:"optional"` // Path relative to the data source root
}

type Entry struct {
	Size  int64  `json:"size"`
	IsDir bool   `json:"isDir"`
	Path  string `json:"path"`
}

type CheckSourceResult struct {
}

// CheckSourceHandler godoc
// @Summary Check the connection of the data source by listing a path
// @Tags Data Source
// @Accept json
// @Produce json
// @Param id path string true "Source ID"
// @Param request body CheckSourceRequest true "Request body"
// @Success 200 {array} Entry
// @Failure 500 {object} handler.HTTPError
// @Router /source/{id}/check [post]
func CheckSourceHandler(
	db *gorm.DB,
	ctx context.Context,
	id string,
	request CheckSourceRequest,
) ([]Entry, *handler.Error) {
	log.SetAllLoggers(log.LevelInfo)
	sourceID, err := strconv.Atoi(id)
	if err != nil {
		return nil, handler.NewBadRequestString("invalid source id")
	}
	var source model.Source
	err = db.Where("id = ?", sourceID).First(&source).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, handler.NewBadRequestString("source not found")
	}
	if err != nil {
		return nil, handler.NewHandlerError(err)
	}

	resolver := datasource.DefaultHandlerResolver{}
	h, err := resolver.Resolve(ctx, source)
	if err != nil {
		return nil, handler.NewHandlerError(err)
	}
	obj, err := h.Check(ctx, request.Path)
	if errors.Is(err, fs.ErrorIsDir) {
		entries, err := h.List(ctx, request.Path)
		if err != nil {
			return nil, handler.NewHandlerError(err)
		}
		return underscore.Map(entries, func(entry fs.DirEntry) Entry {
			_, isDir := entry.(fs.Directory)
			return Entry{
				Size: entry.Size(),
				Path: entry.Remote(),
				IsDir: isDir,
			}
		}), nil
	}

	if err != nil {
		return nil, handler.NewHandlerError(err)
	}

	return []Entry{
		{
			Size: obj.Size(),
			Path: obj.Remote(),
			IsDir: false,
		},
	}, nil
}
