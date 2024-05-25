package service

import (
	model "norht-gowind/src/modules/employees/models"
	repository "norht-gowind/src/modules/employees/repositories"
)

func GetAllEmployees() ([]model.Employee, error) {
	return repository.GetAllEmployees()
}
