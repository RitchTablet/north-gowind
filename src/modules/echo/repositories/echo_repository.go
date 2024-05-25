package repository

import (
	"norht-gowind/src/common/database"
	model "norht-gowind/src/modules/echo/models"
)

func FindAllEchos() ([]model.Echo, error) {
	var echos []model.Echo
	result := database.DB.Find(&echos)
	return echos, result.Error
}
