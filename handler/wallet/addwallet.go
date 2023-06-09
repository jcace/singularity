package wallet

import (
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler"
	"github.com/data-preservation-programs/singularity/model"
	"gorm.io/gorm"
)

// AddWalletHandler godoc
// @Summary Associate a new wallet with a dataset
// @Tags Wallet
// @Produce json
// @Accept json
// @Param name path string true "Dataset name"
// @Param wallet path string true "Wallet Address"
// @Success 200 {object} string
// @Failure 400 {object} handler.HTTPError
// @Failure 500 {object} handler.HTTPError
// @Router /dataset/{datasetName}/wallet/{wallet} [post]
func AddWalletHandler(
	db *gorm.DB,
	datasetName string,
	wallet string,
) (string, *handler.Error) {
	if datasetName == "" {
		return "", handler.NewBadRequestString("dataset name is required")
	}

	if wallet == "" {
		return "", handler.NewBadRequestString("wallet address is required")
	}

	dataset, err := database.FindDatasetByName(db, datasetName)
	if err != nil {
		return "", handler.NewBadRequestString("failed to find dataset: " + err.Error())
	}

	var w model.Wallet
	err = db.Where("address = ?", wallet).First(&w).Error
	if err != nil {
		return "", handler.NewBadRequestString("failed to find wallet: " + err.Error())
	}

	err = database.DoRetry(func() error {
		return db.Create(&model.WalletAssignment{
			DatasetID: dataset.ID,
			WalletID:  w.ID,
		}).Error
	})

	if err != nil {
		return "", handler.NewBadRequestString("failed to create wallet assignment: " + err.Error())
	}

	return wallet, nil
}
