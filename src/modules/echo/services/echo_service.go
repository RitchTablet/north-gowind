package service

import (
	model "norht-gowind/src/modules/echo/models"
	repository "norht-gowind/src/modules/echo/repositories"
)

func GetAllEchos() ([]model.Echo, error) {
	return repository.FindAllEchos()
}
