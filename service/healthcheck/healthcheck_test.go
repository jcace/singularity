package healthcheck

import (
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"testing"
)

func TestHealthCheck(t *testing.T) {
	assert := assert.New(t)
	db := database.OpenInMemory()
	defer database.DropAll(db)

	id := uuid.New()
	HealthCheck(db, id, func() State {
		return State{
			WorkType:  model.Packing,
			WorkingOn: "something",
		}
	})
	var worker model.Worker
	err := db.Where("id = ?", id.String()).First(&worker).Error
	assert.Nil(err)
	assert.Equal(model.Packing, worker.WorkType)
	assert.Equal("something", worker.WorkingOn)
	assert.NotEmpty(worker.Hostname)
	lastHeatbeat := worker.LastHeartbeat

	HealthCheck(db, id, func() State {
		return State{
			WorkType:  model.Packing,
			WorkingOn: "something else",
		}
	})
	err = db.Where("id = ?", id.String()).First(&worker).Error
	assert.Nil(err)
	assert.Equal(model.Packing, worker.WorkType)
	assert.Equal("something else", worker.WorkingOn)
	assert.NotEmpty(worker.Hostname)
	assert.NotEqual(lastHeatbeat, worker.LastHeartbeat)

	HealthCheckCleanup(db)
	err = db.Where("id = ?", id.String()).First(&worker).Error
	assert.Nil(err)

	oldThreshold := staleThreshold
	staleThreshold = 0
	defer func() {
		staleThreshold = oldThreshold
	}()

	HealthCheckCleanup(db)
	err = db.Where("id = ?", id.String()).First(&worker).Error
	assert.ErrorIs(err, gorm.ErrRecordNotFound)
}
