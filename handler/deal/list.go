package deal

import (
	"github.com/data-preservation-programs/singularity/handler"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/rjNemo/underscore"
	"gorm.io/gorm"
)

type ListDealRequest struct {
	Datasets  []string `json:"datasets"`
	Schedules []uint32 `json:"schedules"`
	Providers []string `json:"providers"`
	States    []string `json:"states"`
}

// ListHandler godoc
// @Summary List all deals
// @Description List all deals
// @Tags Deal
// @Accept json
// @Produce json
// @Param request body ListDealRequest true "ListDealRequest"
// @Success 200 {array} model.Deal
// @Failure 400 {object} handler.HTTPError
// @Failure 500 {object} handler.HTTPError
// @Router /deal/list [post]
func ListHandler(db *gorm.DB, request ListDealRequest) ([]model.Deal, *handler.Error) {
	var deals []model.Deal
	statement := db
	if len(request.Datasets) > 0 {
		var datasets []model.Dataset
		err := db.Where("name in ?", request.Datasets).Find(&datasets).Error
		if err != nil {
			return nil, handler.NewHandlerError(err)
		}
		statement = statement.Where("dataset_id IN ?", underscore.Map(datasets, func(dataset model.Dataset) uint32 {return dataset.ID}))
	}

	if len(request.Schedules) > 0 {
		statement = statement.Where("schedule_id IN ?", request.Schedules)
	}

	if len(request.Providers) > 0 {
		statement = statement.Where("provider_id IN ?", request.Providers)
	}

	if len(request.States) > 0 {
		statement = statement.Where("state IN ?", request.States)
	}

	err := db.Find(&deals).Error
	if err != nil {
		return nil, handler.NewHandlerError(err)
	}
	return deals, nil
}
